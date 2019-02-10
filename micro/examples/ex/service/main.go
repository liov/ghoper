package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"micro/protobuf"
	"micro/service/say"
)

func main() {
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)
	service.Init()

	protobuf.RegisterGreeterHandler(service.Server(), new(say.ReSay))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
