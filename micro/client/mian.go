package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"io"
	"log"
	"micro/protobuf"
	"time"
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

	rsp1, err := greeter.SayHelloAgain(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := rsp1.Send(&protobuf.HelloRequest{Name: "John"}); err != nil {
				log.Fatal(err)
			}
			reply, err := rsp1.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetMessage())
			time.Sleep(time.Second)
		}
	}()
	//两边同时recv会阻塞啊，啊啊啊啊
	/*	for {
			reply, err := rsp1.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetMessage())
		}
	*/
	fmt.Println(rsp.GetMessage())
}
