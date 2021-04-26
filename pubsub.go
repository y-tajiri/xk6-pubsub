package pubsub

import (
	"context"

	"github.com/loadimpact/k6/js/modules"
	"github.com/loadimpact/k6/js/common"
	"cloud.google.com/go/pubsub"
)

type PubSub struct {}
type Client struct{
	client *pubsub.Client
	ctxPtr *context.Context
}

func init() {
	modules.Register("k6/x/CloudPubSub", new(PubSub))
}

func (* PubSub) XClient(ctxPtr *context.Context, project string) interface{}{
	rt := common.GetRuntime(*ctxPtr)
	cli, _ := pubsub.NewClient(*ctxPtr, project)
	return common.Bind(rt, &Client{client: cli, ctxPtr: ctxPtr}, ctxPtr)
}

func (cli *Client) Publish(topicNmae, msg string) {
	topic := cli.client.Topic(topicNmae)
	m := &pubsub.Message{
		Data: []byte(msg),
	}
	topic.Publish(*cli.ctxPtr, m)
}