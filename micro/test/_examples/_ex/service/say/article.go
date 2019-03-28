package say

import (
	"context"
	"hoper/test/_examples/_ex/protobuf"
	"io"
)

type Say struct{}

func (g *Say) Hello(stream protobuf.HelloService_HelloServer) error {

	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &protobuf.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

type ReSay struct{}

func (r *ReSay) SayHello(ctx context.Context, in *protobuf.HelloRequest, out *protobuf.HelloReply) error {
	out.Message = "Hello " + in.Name
	return nil
}

func (r *ReSay) SayHelloAgain(ctx context.Context, gstream protobuf.Greeter_SayHelloAgainStream) error {
	for {
		req, err := gstream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &protobuf.HelloReply{Message: "远程调用" + req.GetName()}

		err = gstream.Send(reply)
		if err != nil {
			return err
		}
	}
	return nil
}
