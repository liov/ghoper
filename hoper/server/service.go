package server

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"hoper/protobuf"
	"hoper/server/handler/sub"
	"hoper/server/handler/user"
	"hoper/utils/ulog"
)

func Service() {
	//etcd vs consul
	/*	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs =[]string{
			"http://192.168.3.34:2379",
		}
	})*/
	//consul设置addr
	/*	registry.DefaultRegistry.Init(func(options *registry.Options) {
		options.Addrs = []string{
			"http://192.168.3.34:2379",
		}
	})*/
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Registry(registry.DefaultRegistry),
	)

	// Init will parse the command line flags.
	service.Init()

	protobuf.RegisterUserServiceHandler(service.Server(), new(user.UserHandler))

	// register subscriber
	micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(sub.Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), sub.SubEv, server.SubscriberQueue("queue.pubsub"))

	if err := service.Run(); err != nil {
		ulog.Error(err)
	}
}
