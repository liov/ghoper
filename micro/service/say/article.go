package say

import (
	"context"
	"micro/protobuf"
)

type Say struct{}

func (g *Say) SayHello(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloReply, error) {

	return &protobuf.HelloReply{Message: "Hello " + req.Name}, nil
}

func (g *Say) SayHelloAgain(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloReply, error) {

	return &protobuf.HelloReply{Message: "Hello " + req.Name}, nil
}

type ReSay struct{}

func (r *ReSay) SayHello(ctx context.Context, in *protobuf.HelloRequest, out *protobuf.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}

func (r *ReSay) SayHelloAgain(ctx context.Context, in *protobuf.HelloRequest, out *protobuf.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}
