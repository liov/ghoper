package user

import (
	"context"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
	"micro/common/controller/common"
	"micro/protobuf"
)

var userService protobuf.UserService

func init() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("user.client"))
	service.Init()

	// Create new user client
	userService = protobuf.NewUserService("user", service.Client())
}

func Signup(ctx iris.Context) {

	var user protobuf.SignupReq
	if err := ctx.ReadJSON(&user); err != nil {
		common.Response(ctx, "参数错误")
		return
	}

	rsp, err := userService.Signup(context.TODO(), &user)

	if err != nil {
		common.Response(ctx, err.Error())
		return
	}
	common.Response(ctx, rsp.GetMsg(), rsp.GetUser())
}

func Login(ctx iris.Context) {

	var loginReq protobuf.LoginReq

	if err := ctx.ReadJSON(&loginReq); err != nil {
		common.Response(ctx, "账号或密码错误")
		return
	}

	rsp, err := userService.Login(context.TODO(), &loginReq)

	if err != nil {
		common.Response(ctx, err.Error(), rsp.GetEmail())
		return
	}
	common.Response(ctx, rsp.GetMsg(), rsp.GetUser())
}

func Logout(ctx iris.Context) {
	userInter := ctx.GetViewData()["user"]

	logoutReq := protobuf.LogoutReq{ID: userInter.(protobuf.User).ID}

	rsp, err := userService.Logout(context.TODO(), &logoutReq)

	if err != nil {
		common.Response(ctx, err.Error())
		return
	}
	common.Response(ctx, rsp.GetMsg())
}
