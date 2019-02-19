package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"hoper/protobuf"
	"hoper/service/user/handler"
)

func main() {
	// 创建Service，并定义一些参数
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	service.Init()

	protobuf.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
