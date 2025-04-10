package input

import (
	"context"
	"encore.dev/pubsub"
	"encore.dev/rlog"
)

// encore:api public method=POST path=/send
func Send(ctx context.Context, msg *models.InputMessage) error {
	rlog.Info("Publishing input message", "content", msg.Content)
	_, err := inputTopic.Publish(ctx, *msg)
	return err
}

// Создаем тему с именем "input-messages" и конфигурацией, определяющей гарантию доставки.
var inputTopic = pubsub.NewTopic[models.InputMessage]("input-messages", pubsub.TopicConfig{
	DeliveryGuarantee: pubsub.AtLeastOnce,
})
