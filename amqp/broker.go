package amqp

import (
	"fmt"
	"sync"
	"time"

	"github.com/locngodn/gas-common/util"
	"github.com/streadway/amqp"
)

type (
	RabbitConfig struct {
		Host       string
		Port       int
		User       string
		Password   string
		Ttl        int
		Timeout    int
		MaxRetries int
	}

	AmqpBroker struct {
		sync.Mutex
		amqpClient
		url       string
		conn      *amqp.Connection
		watchStop chan bool
		errors    chan *amqp.Error
	}
)

const (
	XRetries = "x-retries"
)

func getRabbitConnectionString(rabbitConfig *RabbitConfig) string {
	credentials := ""

	if rabbitConfig.User != "" && rabbitConfig.Password != "" {
		credentials = fmt.Sprintf("%s:%s@", rabbitConfig.User, rabbitConfig.Password)
	}

	return fmt.Sprintf("amqp://%s%s:%d/", credentials, rabbitConfig.Host, rabbitConfig.Port)
}

func NewAmqpBroker(rabbitConfig *RabbitConfig, logger util.Logger) *AmqpBroker {
	url := getRabbitConnectionString(rabbitConfig)

	broker := AmqpBroker{
		sync.Mutex{},
		amqpClient{
			amqpPublisherManager: amqpPublisherManager{
				logger:     logger,
				timeout:    rabbitConfig.Timeout,
				maxRetries: rabbitConfig.MaxRetries,
			},
			amqpSubscriptionManager: amqpSubscriptionManager{
				logger: logger,
				ttl:    rabbitConfig.Ttl,
			},
			logger: logger,
		},
		url,
		nil,
		make(chan bool),
		nil,
	}

	return &broker
}

func (b *AmqpBroker) Start() error {
	err := b.setup()
	go b.watch()
	return err
}

func (b *AmqpBroker) Stop() error {
	b.stopWatch()
	defer close(b.watchStop)
	return b.Close()
}

func (b *AmqpBroker) GetLogger() util.Logger {
	return b.logger
}

func (b *AmqpBroker) setup() error {
	b.logger.Info("Setup AMQP Connection")
	if b.conn != nil {
		return nil
	}

	conn, err := b.connect()
	if err != nil {
		b.logger.Errorf("connection failed: %v", err)
		return err
	}

	err = b.Init(conn)
	if err != nil {
		b.logger.Errorf("client init failed: %v", err)
		return err
	}

	b.Lock()
	b.conn = conn
	b.Unlock()

	errors := make(chan *amqp.Error)
	b.conn.NotifyClose(errors)
	b.errors = errors
	return nil
}

func (b *AmqpBroker) connect() (*amqp.Connection, error) {
	return amqp.Dial(b.url)
}

func (b *AmqpBroker) watch() {
	for {
		select {
		case err := <-b.errors:
			b.logger.Warnf("Connection lost: %v\n", err)
			b.Lock()
			b.conn = nil
			b.Unlock()
			time.Sleep(10 * time.Second)
			b.setup()
		case stop := <-b.watchStop:
			if stop {
				return
			}
		}
	}
}

func (b *AmqpBroker) stopWatch() {
	b.watchStop <- true
}
