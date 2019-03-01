package service

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"hoper/client/controller/common/logging"
	"hoper/protobuf"
	"hoper/service/handler/sub"
	"hoper/service/handler/user"
)

func Service() {
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	service.Init()

	protobuf.RegisterUserServiceHandler(service.Server(), new(user.UserHandler))


	// register subscriber
	micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(sub.Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), sub.SubEv, server.SubscriberQueue("queue.pubsub"))

	if err := service.Run(); err != nil {
		logging.Error(err)
	}
}
