package user

import (
	"context"
	"github.com/kataras/iris"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"hoper/client/controller/common"
	"hoper/model/e"
	"hoper/protobuf"
)

var Service protobuf.UserService

func init() {
	app := cmd.App()

	app.Flags = append(app.Flags, cli.StringFlag{
		Name:  "p",
		Usage: "password",
	},
		cli.StringFlag{
			Name:  "mp",
			Usage: "mail password",
		})

	before := app.Before

	app.Before = func(ctx *cli.Context) error {
		if path := ctx.String("c"); len(path) > 0 {
			// got config
			// do stuff
		}
		return before(ctx)
	}
	/*	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs =[]string{
			"http://192.168.3.34:2379",
		}
	})*/
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("user.client"))
	// Init will parse the command line flags.
	//程序自定义命令行，与此冲突 https://github.com/micro/micro/issues/83
	service.Init()
	// Create new user client
	Service = protobuf.NewUserService("user", service.Client())
}

func Signup(ctx iris.Context) {

	var user protobuf.SignupReq
	if err := ctx.ReadJSON(&user); err != nil {
		common.Response(ctx, nil, "参数错误", e.ERROR)
		return
	}

	rsp, err := Service.Signup(context.TODO(), &user)

	if err != nil {
		common.Response(ctx, nil, err.Error(), e.ERROR)
		return
	}
	common.Response(ctx, rsp.GetUser(), rsp.GetMsg(), e.SUCCESS)
}

func Login(ctx iris.Context) {

	var loginReq protobuf.LoginReq

	if err := ctx.ReadJSON(&loginReq); err != nil {
		common.Response(ctx, nil, "账号或密码错误", e.ERROR)
		return
	}

	rsp, err := Service.Login(context.TODO(), &loginReq)

	if err != nil {
		common.Response(ctx, nil, err.Error(), e.ERROR)
		return
	}
	common.Response(ctx, rsp.GetUser(), rsp.GetMsg(), e.SUCCESS)
}

func Logout(ctx iris.Context) {
	userInter := ctx.Values().Get("user")

	logoutReq := protobuf.LogoutReq{ID: userInter.(protobuf.User).ID}

	rsp, err := Service.Logout(context.TODO(), &logoutReq)

	if err != nil {
		common.Response(ctx, nil, err.Error(), e.ERROR)
		return
	}
	common.Response(ctx, nil, rsp.GetMsg(), e.SUCCESS)
}

func GetUser(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	getReq := protobuf.GetReq{ID: id}
	user, err := Service.GetUser(context.TODO(), &getReq)
	if err != nil {
		common.Response(ctx, nil, err.Error(), e.ERROR)
		return
	}
	common.Response(ctx, user, e.GetMsg(e.SUCCESS), e.SUCCESS)
}
