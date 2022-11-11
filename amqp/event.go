package amqp

import (
	"time"

	"github.com/locngodn/gas-common/constant"
	"github.com/locngodn/gas-common/model"
	"github.com/streadway/amqp"
)

type (
	amqpEvent struct {
		delivery amqp.Delivery
	}
)

func (e *amqpEvent) GetHeaders() amqp.Table {
	return e.delivery.Headers
}

func (e *amqpEvent) GetStatus() *model.Status {
	headers := e.GetHeaders()
	if headers[constant.StatusCode] != nil {
		statusCode := int(headers[constant.StatusCode].(int32))
		statusText := headers[constant.StatusText].(string)
		return &model.Status{StatusCode: statusCode, StatusText: statusText}
	}
	return nil
}

func (e *amqpEvent) IsSuccess() bool {
	if e.GetStatus() == nil || (e.GetStatus().StatusCode >= 200 && e.GetStatus().StatusCode <= 300) {
		return true
	}
	return false
}

func (e *amqpEvent) GetLanguage() string {
	headers := e.GetHeaders()
	if headers[constant.Language] != nil {
		return headers[constant.Language].(string)
	}
	return ""
}

func (e *amqpEvent) GetCorrelationID() string {
	return e.delivery.CorrelationId
}

func (e *amqpEvent) GetReplyTo() string {
	return e.delivery.ReplyTo
}

func (e *amqpEvent) GetBody() []byte {
	return e.delivery.Body
}

func (e *amqpEvent) GetExpiration() time.Time {
	return e.delivery.Timestamp
}
func (e *amqpEvent) GetUserId() string {
	return e.delivery.UserId
}
func (e *amqpEvent) GetMessageId() string {
	return e.delivery.MessageId
}

func (e *amqpEvent) Ack() error {
	return e.delivery.Ack(false)
}

func (e *amqpEvent) Nack(requeue bool) error {
	return e.delivery.Nack(false, requeue)
}

func (e *amqpEvent) Reject(requeue bool) error {
	return e.delivery.Reject(requeue)
}
