package main

import (
	"github.com/micro/go-micro"
	"micro/protobuf"
	"micro/service/say"
)

func main() {
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	service.Init()

	protobuf.RegisterGreeterHandler(service.Server(), new(say.ReSay))

	service.Run()
}
