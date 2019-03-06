package user

import (
	"context"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/protobuf"
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
	common.Response(ctx, rsp.GetUser(), rsp.GetMsg())
}

func Login(ctx iris.Context) {

	var loginReq protobuf.LoginReq

	if err := ctx.ReadJSON(&loginReq); err != nil {
		common.Response(ctx, "账号或密码错误")
		return
	}

	rsp, err := userService.Login(context.TODO(), &loginReq)

	if err != nil {
		common.Response(ctx, err.Error())
		return
	}
	common.Response(ctx, rsp.GetUser(), rsp.GetMsg())
}

func Logout(ctx iris.Context) {
	userInter := ctx.Values().Get("user")

	logoutReq := protobuf.LogoutReq{ID: userInter.(protobuf.User).ID}

	rsp, err := userService.Logout(context.TODO(), &logoutReq)

	if err != nil {
		common.Response(ctx, err.Error())
		return
	}
	common.Response(ctx, rsp.GetMsg())
}

func GetUser(ctx iris.Context) {
	id := ctx.Params().GetUint64Default("id", 0)
	getReq := protobuf.GetReq{ID: id}
	user, err := userService.GetUser(context.TODO(), &getReq)
	if err != nil {
		common.Response(ctx, err.Error())
		return
	}
	common.Response(ctx, user, e.GetMsg(e.SUCCESS), e.SUCCESS)
}
