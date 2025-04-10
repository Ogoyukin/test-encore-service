package process

import (
	"context"
	"encore.dev/pubsub"
	"encore.dev/rlog"
)

// encore:topic name=input-messages
var _ = pubsub.Subscribe(inputHandler)

func inputHandler(ctx context.Context, msg *models.InputMessage) error {
	reversed := reverse(msg.Content)
	rlog.Info("Processing message", "original", msg.Content, "reversed", reversed)
	return outputTopic.Publish(ctx, &models.InputMessage{Content: reversed})
}

// encore:topic name=processed.messages
var outputTopic pubsub.Topic[models.InputMessage]

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
