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

func init() {
	modules.Register("k6/x/CloudPubSub", new(PubSub))
}

func (* PubSub) XClient(ctxPtr *context.Context) interface{}{
	rt := common.GetRuntime(*ctxPtr)
	//cli, _ := pubsub.NewClient(*ctxPtr, project)
	return common.Bind(rt, &Client{}, ctxPtr)
}
