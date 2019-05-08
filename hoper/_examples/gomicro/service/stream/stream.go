package main

import (
	"google.golang.org/grpc"
	"hoper/_examples/gomicro/protobuf"
	"hoper/_examples/gomicro/service/say"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	protobuf.RegisterHelloServiceServer(grpcServer, new(say.Say))

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
