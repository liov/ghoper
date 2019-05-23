package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
	"hoper/_examples/gomicro/protobuf"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := protobuf.NewGreeterService("greeter", service.Client())

	rsp1, err := greeter.SayHelloAgain(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	app := iris.New()

	v1 := app.Party("/iris")
	{
		v1.Get("/{msg:string}", func(ctx iris.Context) {

			msg := ctx.Params().GetString("msg")

			// Call the greeter
			rsp, err := greeter.SayHello(context.TODO(), &protobuf.HelloRequest{Name: msg})
			if err != nil {
				ulog.Error(err)
			}

			ctx.JSON(iris.Map{
				"message": rsp.GetMessage(),
			})

		})
	}

	// listen and serve on http://0.0.0.0:8080.
	go func() {
		// listen and serve on http://0.0.0.0:8080.
		if err := app.Run(iris.Addr(":8888")); err != nil && err != http.ErrServerClosed {
			log.Printf("Listen: %s\n", err)
		}
	}()

	go func() {
		for {
			if err := rsp1.Send(&protobuf.HelloRequest{Name: "grpc流"}); err != nil {
				log.Fatal(err)
			}

			reply, _ := rsp1.Recv()

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
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
