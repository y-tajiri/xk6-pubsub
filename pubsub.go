package pubsub

import (
	"context"

	"github.com/loadimpact/k6/js/modules"
	"cloud.google.com/go/pubsub"
)

type PubSub struct {}

func init() {
	modules.Register("k6/x/CloudPubSub", new(PubSub))
}

func (* PubSub) NewClient(ctx context.Context, project string) *pubsub.Client {
	cli, _ := pubsub.NewClient(ctx, project)
	return cli
}