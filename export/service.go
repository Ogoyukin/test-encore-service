package export

import (
	"context"
	"encore.dev/pubsub"
	"encore.dev/rlog"
	"fmt"
)

// encore:topic name=processed.messages
var _ = pubsub.Subscribe(printHandler)

func printHandler(ctx context.Context, msg *models.InputMessage) error {
	rlog.Info("Exporting message", "content", msg.Content)
	for i := 0; i < 3; i++ {
		fmt.Println("[export-service]", msg.Content)
	}
	return nil
}
