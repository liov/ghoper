package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
	"log"
	"micro/protobuf"
	"net/http"
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

	app := iris.New()

	v1 := app.Party("/iris")
	{
		v1.Get("/{msg:string}", func(ctx iris.Context) {

			msg := ctx.Params().GetString("msg")

			if err := rsp1.Send(&protobuf.HelloRequest{Name: msg}); err != nil {
				log.Fatal(err)
			}

			reply, _ := rsp1.Recv()

			ctx.JSON(iris.Map{
				"message": reply.GetMessage(),
			})

		})
	}

	// listen and serve on http://0.0.0.0:8080.
	if err := app.Run(iris.Addr(":8888")); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}
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
