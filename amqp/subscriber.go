package amqp

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/locngodn/gas-common/constant"
	"github.com/locngodn/gas-common/messaging"
	"github.com/locngodn/gas-common/model"
	"github.com/locngodn/gas-common/util"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func getConsumerID(queue string) string {
	var hostname string

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}

	pid := os.Getpid()

	return fmt.Sprintf("%s#%d@%s/%d", queue, pid, hostname, rand.Uint32())
}

func newAmqpSubscription(queueName, key, exchangeType string, exchange string, autoAck bool, autoDelete bool, exclusive bool, workerCount int, ttl int, processor messaging.SubscribeProcessor, logger util.Logger) *amqpSubscriber {
	return &amqpSubscriber{
		QueueName:    queueName,
		Key:          key,
		Exchange:     exchange,
		ExchangeType: exchangeType,
		AutoAck:      autoAck,
		AutoDelete:   autoDelete,
		Exclusive:    exclusive,
		workerCount:  workerCount,
		processor:    processor,
		ttl:          ttl,
		Logger:       logger.WithFields(util.BuildContext("amqpSubscription")).WithField("queueName", queueName).WithField("exchangeName", exchange),
	}
}

type amqpSubscriptionManager struct {
	connection  *amqp.Connection
	ttl         int
	logger      util.Logger
	Subscribers []*amqpSubscriber
}

func (sub *amqpSubscriptionManager) Init(connection *amqp.Connection, ttl int, logger util.Logger) error {
	sub.connection = connection
	sub.ttl = ttl
	sub.logger = logger

	if sub.Subscribers != nil && len(sub.Subscribers) > 0 {
		for _, subscriber := range sub.Subscribers {
			subscriber.init(connection)
		}
	}
	return nil
}

func (sub *amqpSubscriptionManager) CreateSubscription(queueName, key, exchangeType string, exchange string, autoAck bool, autoDelete bool, exclusive bool, workerCount int, ttl int, processor messaging.SubscribeProcessor) (messaging.Subscription, error) {
	var t int
	if ttl > 0 {
		t = ttl
	} else {
		t = sub.ttl
	}
	subscription := newAmqpSubscription(queueName, key, exchangeType, exchange, autoAck, autoDelete, exclusive, workerCount, t, processor, sub.logger)
	if err := subscription.init(sub.connection); err != nil {
		return nil, err
	}
	sub.Subscribers = append(sub.Subscribers, subscription)
	return subscription, nil
}

func (client *amqpSubscriptionManager) Close() error {
	var err error
	for _, sub := range client.Subscribers {
		err = sub.Close()
	}
	return err
}

type amqpSubscriber struct {
	AutoAck      bool
	AutoDelete   bool
	Exclusive    bool
	QueueName    string
	Key          string
	ExchangeType string
	Exchange     string
	consumerId   string
	workerCount  int
	processingWG sync.WaitGroup // use wait group to make sure task processing completes on interrupt signal
	processor    messaging.SubscribeProcessor
	stopChan     chan int
	channel      *amqp.Channel
	ttl          int
	Logger       util.Logger
}

func (c *amqpSubscriber) init(connection *amqp.Connection) error {
	c.consumerId = getConsumerID(c.QueueName)
	c.Logger.Infof("initiated")

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	c.stopChan = make(chan int)
	errorChan := channel.NotifyClose(make(chan *amqp.Error))

	go func() {
		// NOTE: We need this in case the channel gets broken (but not the connection), e.g.
		// when an invalid message is acknowledged.
		for err := range errorChan {
			c.Logger.Errorf("queue-channel error %v\n", err)
			c.channel.Cancel(c.consumerId, false)
			c.stopChan <- 1
			//if cerr := c.init(connection); cerr != nil {
			//	util.Log.Errorf("Failed to reestablish channel for queue %s: %v\n", c.QueueName, cerr)
			//}
		}
	}()

	if len(c.ExchangeType) > 0 && len(c.Exchange) > 0 {
		err = channel.ExchangeDeclare(
			c.Exchange,     // name
			c.ExchangeType, // type
			true,           // durable
			false,          // audo-delete
			false,          // internal
			false,          // no-wait
			nil,            // args
		)
		if err != nil {
			return err
		}
	}

	if err := channel.Qos(c.workerCount*2, 0, false); err != nil {
		return err
	}

	table := make(amqp.Table)
	deadLetterQueue := c.QueueName + constant.DEAD_LETTER
	table["x-message-ttl"] = c.ttl       // in ms
	table["x-dead-letter-exchange"] = "" //c.Exchange
	table["x-dead-letter-routing-key"] = c.QueueName
	if c.AutoDelete {
		table["x-expires"] = 1000
	}

	if _, err := channel.QueueDeclare(deadLetterQueue, true, c.AutoDelete, c.Exclusive, false, table); err != nil {
		return err
	}

	table = make(amqp.Table)
	table["x-dead-letter-exchange"] = "" //c.Exchange
	table["x-dead-letter-routing-key"] = deadLetterQueue
	if _, err := channel.QueueDeclare(c.QueueName, true, c.AutoDelete, c.Exclusive, false, table); err != nil {
		return err
	}

	if c.Exchange != "" {
		if err := channel.QueueBind(c.QueueName, c.Key, c.Exchange, false, nil); err != nil {
			return err
		}
	}

	deliveryChan, err := channel.Consume(c.QueueName, c.consumerId, c.AutoAck, false, false, false, nil)
	if err != nil {
		return err
	}

	c.channel = channel

	go c.runEventProcessor(deliveryChan)

	c.Logger.Infof("initiated done")
	return nil
}

func (c *amqpSubscriber) runEventProcessor(deliveryChan <-chan amqp.Delivery) {
	pool := util.NewIntPool(c.workerCount)

	for {
		select {
		case d := <-deliveryChan:
			if len(d.Body) == 0 {
				d.Nack(false, false)
				return
			}

			c.Logger.Infof("recieved message %s", d.CorrelationId)

			if !d.Timestamp.IsZero() && d.Timestamp.Before(time.Now()) {
				c.Logger.Infof("Message Expired", d.CorrelationId)
				d.Ack(false)
			} else {

				workerNumber := pool.Get()
				event := amqpEvent{d}
				c.processingWG.Add(1)

				go func(worker int, event *amqpEvent, log util.Logger) {
					beginTime := time.Now()
					logger := log.WithField("workerNumber", worker).
						WithField("correlationId", event.GetCorrelationID())

					ctx := util.NewSessionCtx(event.GetCorrelationID(), util.Logger(logger))

					logger.WithField("message", string(event.GetBody())).Info("process event started")
					err := c.processor(ctx, event)
					logger.Info("process event finish")

					defer func() {
						c.processingWG.Done()
						pool.Put(worker)
					}()

					defer func() {
						status := model.NewStatus(err, false)
						duration := time.Since(beginTime)
						log := logger.WithFields(logrus.Fields{
							"duration":   duration,
							"statusCode": status.StatusCode,
							"statusText": status.StatusText,
						})
						log.Infof("")
					}()

				}(workerNumber, &event, c.Logger)
			}
		case <-c.stopChan:
			c.Logger.Debugf("stop triggered")
			return
		}
	}
}

func (c *amqpSubscriber) Close() error {
	c.Logger.Debugf("close triggered")
	close(c.stopChan)

	// Waiting for any tasks being processed to finish
	c.processingWG.Wait()

	// NOTE: This should close the deliveryChannel, which quits the loop in Run(), which stops this subscriber
	if err := c.channel.Cancel(c.consumerId, false); err != nil {
		return err
	}
	return c.channel.Close()
}
