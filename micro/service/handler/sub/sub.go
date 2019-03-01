package sub

import (
	"context"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/metadata"
	"hoper/protobuf"
)

type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *protobuf.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func SubEv(ctx context.Context, event *protobuf.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
