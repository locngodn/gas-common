package messaging

import (
	"context"
	"time"

	"github.com/locngodn/gas-common/model"
	"github.com/locngodn/gas-common/util"
	"github.com/streadway/amqp"

	"github.com/gogo/protobuf/proto"
)

type (
	OptionMessage struct {
		CorrelationId string
		ReplyTo       string
		Language      string
		UserId        string
		MessageId     string
		Expiration    time.Time
	}

	PublishProcessor = func(event Event) error

	Publisher interface {
		Init(option map[string]interface{}) error
		Publish(err error, payload proto.Message, options ...OptionMessage) error
		BasicPublish(payload proto.Message, options ...OptionMessage) error
		BroadcastPublish(exchangeName string, payload proto.Message, options ...OptionMessage) error
		PublishRpc(payload proto.Message, handler PublishProcessor, options ...OptionMessage) error
		PublishRequeue(event Event) error
	}

	Event interface {
		GetHeaders() amqp.Table
		GetStatus() *model.Status
		GetLanguage() string
		GetCorrelationID() string
		GetReplyTo() string
		GetBody() []byte
		GetExpiration() time.Time
		GetUserId() string
		GetMessageId() string
		IsSuccess() bool
		Ack() error
		Nack(requeue bool) error
		Reject(requeue bool) error
	}

	Subscription interface {
		Close() error
	}

	SubscribeProcessor = func(ctx context.Context, event Event) error

	Broker interface {
		Start() error
		Stop() error
		GetLogger() util.Logger
		CreateSubscription(queueName, key, exchangeType string, exchange string, autoAck bool, autoDelete bool, exclusive bool, workerCount int, ttl int, processor SubscribeProcessor) (Subscription, error)
		CreatePublisher(publishKey string) (Publisher, error)
	}
)

func (opt OptionMessage) Merge(m OptionMessage) {
	if len(opt.CorrelationId) == 0 {
		opt.CorrelationId = m.CorrelationId
	}
	if len(opt.ReplyTo) == 0 {
		opt.ReplyTo = m.ReplyTo
	}
	if len(opt.Language) == 0 {
		opt.Language = m.Language
	}
	if len(opt.UserId) == 0 {
		opt.UserId = m.UserId
	}
	if len(opt.MessageId) == 0 {
		opt.MessageId = m.MessageId
	}
	if opt.Expiration.IsZero() {
		opt.Expiration = m.Expiration
	}
}
