package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"service/controller"
	"service/controller/common"
	"service/controller/common/e"
	"service/initialize"
	"service/model"
)

func JWT(ctx iris.Context) {
	code := e.SUCCESS
	user, err := getUser(ctx)
	if err != nil && err.Error() == "没有token" {
		code = e.InvalidParams
	} else if err != nil && err.Error() == "未登录" {
		code = e.ErrorAuthCheckTokenFail
	} else if err != nil && err.Error() == "登录超时" {
		code = e.ErrorAuthCheckTokenTimeout
	}

	if code != e.SUCCESS {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"code": code,
			"data": e.GetMsg(code)})
		return
	}
	ctx.ViewData("user", user)
	ctx.Next()
}

func GetUser() iris.Handler {
	return func(ctx iris.Context) {
		user, _ := getUser(ctx)
		ctx.ViewData("user", user)
		ctx.Next()
	}
}

func getUser(ctx iris.Context) (controller.User, error) {
	var user controller.User
	tokenString := ctx.GetHeader("token")
	if len(tokenString) == 0 {
		tokenString = ctx.GetHeader("Authorization")
	}
	if len(tokenString) == 0 {
		return user, errors.New("没有token")
	}

	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected login method: %v", token.Header["alg"])
		}
		return []byte(initialize.Config.Server.TokenSecret), nil
	})

	if tokenErr != nil {
		return user, errors.New("未登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		var err error
		user, err = controller.UserFromRedis(userID)
		if err != nil {
			return user, errors.New("登录超时")
		}
		return user, nil
	}
	return user, errors.New("未登录")
}

// SetContextUser 给 context 设置 user
func SetContextUser(ctx iris.Context) {
	var user controller.User
	var err error
	if user, err = getUser(ctx); err != nil {
		ctx.ViewData("user", nil)
		return
	}
	ctx.ViewData("user", user)
}

// SigninRequired 必须是登录用户
func SigninRequired(ctx iris.Context) {

	var user controller.User
	var err error
	if user, err = getUser(ctx); err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	ctx.ViewData("user", user)
}

// EditorRequired 必须是网站编辑
func EditorRequired(ctx iris.Context) {

	var user controller.User
	var err error
	if user, err = getUser(ctx); err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleEditor || user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		ctx.ViewData("user", user)
	} else {
		common.Response(ctx, "没有权限", 1003)
	}
}

// AdminRequired 必须是管理员
func AdminRequired(ctx iris.Context) {

	var user controller.User
	var err error
	if user, err = getUser(ctx); err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		ctx.ViewData("user", user)
	} else {
		common.Response(ctx, "没有权限", 1003)
	}
}
