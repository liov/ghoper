package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hoper/_examples/gomicro/protobuf"
	"io"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := protobuf.NewHelloServiceClient(conn)
	stream, err := client.Hello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&protobuf.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
