package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"micro/protobuf"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := protobuf.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.SayHello(context.TODO(), &protobuf.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.GetMessage())
}
