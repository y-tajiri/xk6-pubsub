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
}

type Topic struct{
	topic *pubsub.Topic
}

func init() {
	modules.Register("k6/x/CloudPubSub", new(PubSub))
}

func (* PubSub) XClient(ctxPtr *context.Context, project string) interface{}{
	rt := common.GetRuntime(*ctxPtr)
	cli, _ := pubsub.NewClient(*ctxPtr, project)
	return common.Bind(rt, &Client{client: cli}, ctxPtr)
}

func (cli *Client) XTopic(ctxPtr *context.Context, topicName string) interface{} {
	rt := common.GetRuntime(*ctxPtr)
	topic := cli.client.Topic(topicName)
	return common.Bind(rt, &Topic{topic: topic}, ctxPtr)
}

func (t *Topic)XPublish(ctxPtr *context.Context, msg string){
	m := &pubsub.Message{
		Data: []byte(msg),
	}
	t.topic.Publish(*ctxPtr, m)
}