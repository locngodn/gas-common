package amqp

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/locngodn/gas-common/constant"
	"github.com/locngodn/gas-common/messaging"
	"github.com/locngodn/gas-common/model"
	"github.com/locngodn/gas-common/util"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type (
	amqpPublisherManager struct {
		connection *amqp.Connection
		Publishers map[string]*amqpPublisher
		timeout    int
		maxRetries int
		logger     util.Logger
	}

	amqpPublisher struct {
		publishKey     string
		publishChannel *amqp.Channel
		timeout        int
		maxRetries     int
		logger         util.Logger
	}
)

const (
	messagingContentType = "application/json"
)

func (pub *amqpPublisherManager) Init(connection *amqp.Connection, timeout int, maxRetries int, logger util.Logger) error {
	pub.connection = connection
	pub.timeout = timeout
	pub.maxRetries = maxRetries
	var err error
	if pub.Publishers != nil && len(pub.Publishers) > 0 {
		for _, publisher := range pub.Publishers {
			err = publisher.init(connection)
		}
	}
	return err
}

func (pub *amqpPublisherManager) CreatePublisher(publishKey string) (messaging.Publisher, error) {
	if len(pub.Publishers) > 0 && pub.Publishers[publishKey] != nil {
		return pub.Publishers[publishKey], nil
	}

	log := pub.logger.WithField("publishKey", publishKey)
	log.Info("create publisher")
	publisher := amqpPublisher{
		publishKey: publishKey,
		timeout:    pub.timeout,
		maxRetries: pub.maxRetries,
		logger:     log,
	}

	if err := publisher.init(pub.connection); err != nil {
		return nil, err
	}

	if len(pub.Publishers) == 0 {
		pub.Publishers = make(map[string]*amqpPublisher)
	}

	if !strings.HasPrefix(publishKey, constant.RandomQueuePrefix) {
		pub.Publishers[publishKey] = &publisher
	}

	return &publisher, nil
}

func (pub *amqpPublisherManager) Close() error {
	var err error
	for _, publisher := range pub.Publishers {
		err = publisher.Close()
	}
	return err
}

func (client *amqpPublisher) init(connection *amqp.Connection) error {
	client.logger.Infof("Amqp publisher initiated %s", client.publishKey)
	channel, err := connection.Channel()
	if err != nil {
		return err
	}
	client.publishChannel = channel

	return nil
}

func (client *amqpPublisher) PublishRpc(payload proto.Message, handler messaging.PublishProcessor, options ...messaging.OptionMessage) error {
	if client.publishChannel == nil {
		return fmt.Errorf("Connection currently not available.")
	}

	option := messaging.OptionMessage{}
	if len(options) > 0 {
		option = options[0]
	}
	if len(option.CorrelationId) == 0 {
		option.CorrelationId = uuid.New().String()
	}
	log := client.logger.WithFields(logrus.Fields{
		"correlationId": option.CorrelationId,
		"messageId":     option.MessageId,
	})
	if handler != nil {
		q, err := client.createReplyQueue(option.ReplyTo)
		if err != nil {
			return err
		}
		option.ReplyTo = q.Name
	}

	exchangeName := ""
	routingKey := client.publishKey

	data, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	publishing := amqp.Publishing{
		//ContentType:     messagingContentType,
		ContentEncoding: "UTF-8",
		CorrelationId:   option.CorrelationId,
		ReplyTo:         option.ReplyTo,
		MessageId:       option.MessageId,
		//UserId:          option.UserId,
		Timestamp:    option.Expiration,
		DeliveryMode: amqp.Persistent,
		Body:         data,
	}

	err = client.publishChannel.Publish(exchangeName, routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	log.WithField("option", option).Info(payload)
	if handler != nil {
		msgs, err := client.createConsume(option.ReplyTo)
		if err != nil {
			return err
		}

		var tout time.Duration
		if option.Expiration.IsZero() {
			tout = time.Duration(client.timeout) * time.Millisecond
		} else {
			tout = option.Expiration.Sub(time.Now())
		}
		timeout := make(chan bool, 1)
		go func() {
			select {
			case <-timeout:
			case <-time.After(tout):
				publishing.Body = nil
				headers := make(map[string]interface{})
				headers[constant.StatusCode] = http.StatusRequestTimeout
				headers[constant.StatusText] = constant.RequestTimeout
				publishing.Headers = headers
				client.publishChannel.Publish(exchangeName, option.ReplyTo, false, false, publishing)
			}
		}()

		for d := range msgs {
			if option.CorrelationId == d.CorrelationId {
				event := amqpEvent{d}
				go func(event *amqpEvent, queueName string) {
					handler(event)
					defer func() {
						client.publishChannel.Cancel(d.ConsumerTag, true)
						timeout <- true
					}()
				}(&event, option.ReplyTo)
				break
			}
		}
	}
	return nil
}

func (client *amqpPublisher) BasicPublish(payload proto.Message, options ...messaging.OptionMessage) error {
	if client.publishChannel == nil {
		return fmt.Errorf("Connection currently not available.")
	}
	option := messaging.OptionMessage{}
	if len(options) > 0 {
		option = options[0]
	}
	if len(option.CorrelationId) == 0 {
		option.CorrelationId = uuid.New().String()
	}
	log := client.logger.WithFields(logrus.Fields{
		"correlationId": option.CorrelationId,
		"messageId":     option.MessageId,
	})
	exchangeName := ""
	routingKey := client.publishKey

	data, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	headers := make(map[string]interface{})
	headers[constant.Language] = option.Language

	publishing := amqp.Publishing{
		//ContentType:     messagingContentType,
		Headers:         headers,
		ContentEncoding: "UTF-8",
		CorrelationId:   option.CorrelationId,
		MessageId:       option.MessageId,
		//UserId:          option.UserId,
		Timestamp:    option.Expiration,
		DeliveryMode: amqp.Persistent,
		Body:         data,
	}

	err = client.publishChannel.Publish(exchangeName, routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	log.WithField("option", option).Info(payload)
	defer func() {
		if strings.HasPrefix(routingKey, constant.RandomQueuePrefix) {
			client.Close()
		}
	}()
	return nil
}

func (client *amqpPublisher) BroadcastPublish(exchangeName string, payload proto.Message, options ...messaging.OptionMessage) error {
	if client.publishChannel == nil {
		return fmt.Errorf("Connection currently not available.")
	}
	option := messaging.OptionMessage{}
	if len(options) > 0 {
		option = options[0]
	}
	if len(option.CorrelationId) == 0 {
		option.CorrelationId = uuid.New().String()
	}
	log := client.logger.WithFields(logrus.Fields{
		"correlationId": option.CorrelationId,
		"messageId":     option.MessageId,
		"exchangeName":  exchangeName,
	})
	routingKey := client.publishKey

	data, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	publishing := amqp.Publishing{
		//ContentType:     messagingContentType,
		ContentEncoding: "UTF-8",
		CorrelationId:   option.CorrelationId,
		MessageId:       option.MessageId,
		//UserId:          option.UserId,
		Timestamp:    option.Expiration,
		DeliveryMode: amqp.Persistent,
		Body:         data,
	}

	err = client.publishChannel.Publish(exchangeName, routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	log.WithField("option", option).Info(payload)
	return nil
}

func (client *amqpPublisher) Publish(err error, payload proto.Message, options ...messaging.OptionMessage) error {
	if client.publishChannel == nil {
		return fmt.Errorf("Connection currently not available.")
	}
	option := messaging.OptionMessage{}
	if len(options) > 0 {
		option = options[0]
	}
	if len(option.CorrelationId) == 0 {
		option.CorrelationId = uuid.New().String()
	}
	log := client.logger.WithFields(logrus.Fields{
		"correlationId": option.CorrelationId,
		"messageId":     option.MessageId,
	})
	exchangeName := ""
	routingKey := client.publishKey

	var data []byte
	hasPayload := payload != nil && !reflect.ValueOf(payload).IsNil()
	if hasPayload {
		data, err = proto.Marshal(payload)
		if err != nil {
			return err
		}
	}

	headers := make(map[string]interface{})

	status := model.NewStatus(err, !hasPayload)
	headers[constant.StatusCode] = status.StatusCode
	headers[constant.StatusText] = status.StatusText
	headers[constant.Language] = option.Language

	publishing := amqp.Publishing{
		//ContentType:     messagingContentType,
		Headers:         headers,
		ContentEncoding: "UTF-8",
		CorrelationId:   option.CorrelationId,
		MessageId:       option.MessageId,
		//UserId:          option.UserId,
		Timestamp:    option.Expiration,
		DeliveryMode: amqp.Persistent,
		Body:         data,
	}

	err = client.publishChannel.Publish(exchangeName, routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	log.WithField("option", option).Info(payload)
	defer func() {
		if strings.HasPrefix(routingKey, constant.RandomQueuePrefix) {
			client.Close()
		}
	}()

	return nil
}

func (client *amqpPublisher) PublishRequeue(event messaging.Event) error {
	if client.publishChannel == nil {
		return fmt.Errorf("Connection currently not available.")
	}

	log := client.logger.WithFields(logrus.Fields{
		"correlationId": event.GetCorrelationID(),
		"messageId":     event.GetMessageId(),
	})
	headers := event.GetHeaders()
	val, ok := headers[XRetries]
	var count int32
	if ok {
		count, ok = val.(int32)
		if ok {
			count += 1
		} else {
			count = 0
		}
	} else {
		count = 0
	}

	if client.maxRetries > 0 && int(count) >= client.maxRetries {
		log.Info("Max Requeued")
		return nil
	}

	headers = make(map[string]interface{})
	headers[XRetries] = count

	exchangeName := ""
	routingKey := client.publishKey

	publishing := amqp.Publishing{
		//ContentType:     messagingContentType,
		Headers:         headers,
		ContentEncoding: "UTF-8",
		CorrelationId:   event.GetCorrelationID(),
		ReplyTo:         event.GetReplyTo(),
		MessageId:       event.GetMessageId(),
		//UserId:          event.GetUserId(),
		Timestamp:    event.GetExpiration(),
		DeliveryMode: amqp.Persistent,
		Body:         event.GetBody(),
	}

	err := client.publishChannel.Publish(exchangeName, routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	log.WithField("retries", count).Info(string(event.GetBody()))
	defer func() {
		if strings.HasPrefix(routingKey, constant.RandomQueuePrefix) {
			client.Close()
		}
	}()
	return nil
}

func (client *amqpPublisher) Init(option map[string]interface{}) error {
	if _, err := client.publishChannel.QueueDeclare(client.publishKey, true, false, false, false, option); err != nil {
		return err
	}

	return nil
}

func (client *amqpPublisher) createReplyQueue(name string) (amqp.Queue, error) {
	return client.publishChannel.QueueDeclare(name, true, true, false, false, nil)
}

func (client *amqpPublisher) createConsume(queue string) (<-chan amqp.Delivery, error) {
	return client.publishChannel.Consume(queue, "", true, false, false, false, nil)
}

func (client *amqpPublisher) Close() error {
	return client.publishChannel.Close()
}

func (client *amqpPublisher) QueueDelete(queueName string) error {
	_, err := client.publishChannel.QueueDelete(queueName, true, true, false)
	return err
}
