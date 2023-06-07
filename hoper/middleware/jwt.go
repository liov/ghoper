package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/kataras/iris"
	"hoper/controller"
	"hoper/controller/common"
	"hoper/model"
	"hoper/model/e"
)

// 中间件的两种方式
func GetUser(fullInfo bool) iris.Handler {
	return func(ctx iris.Context) {
		code := e.SUCCESS
		var user *controller.User
		var userID uint64
		var err error
		if fullInfo {
			user, err = getUser(ctx)
		} else {
			userID, err = getUserID(ctx)
		}

		if err != nil && err.Error() == "未登录" {
			code = e.ErrorAuthCheckTokenFail
		} else if err != nil && err.Error() == "登录超时" {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			//ctx.StatusCode(iris.StatusUnauthorized)
			ctx.SetCookie(&http.Cookie{
				Name:     "token",
				Value:    "del",
				Path:     "/",
				Domain:   ctx.Host(),
				Expires:  time.Now().Add(-1),
				MaxAge:   -1,
				Secure:   false,
				HttpOnly: true,
			})
			common.Response(ctx, nil, err.Error(), code)
			return
		}

		if fullInfo {
			ctx.Values().Set("user", user) //指针
		}
		ctx.Values().Set("userID", userID)
		ctx.Next()
	}
}

func GetUserId(ctx iris.Context) {
	if userID, err := getUserID(ctx); err == nil {
		ctx.Values().Set("userID", userID)
	} else {
		ctx.Values().Set("userID", uint64(0))
	}
	ctx.Next()
}

func getUser(ctx iris.Context) (*controller.User, error) {

	tokenString := ctx.GetCookie("token")
	if len(tokenString) == 0 {
		tokenString = ctx.GetHeader("Authorization")
	}
	if len(tokenString) == 0 {
		return nil, errors.New("未登录")
	}

	claims, err := controller.ParseToken(tokenString)

	if err != nil {
		return nil, err
	}
	user, err := controller.UserFromRedis(claims.UserID)

	controller.UserLastActiveTime(claims.UserID)

	return user, nil
}

func getUserID(ctx iris.Context) (uint64, error) {
	tokenString := ctx.GetCookie("token")
	if len(tokenString) == 0 {
		tokenString = ctx.GetHeader("Authorization")
	}
	if len(tokenString) == 0 {
		return 0, errors.New("未登录")
	}

	claims, err := controller.ParseToken(tokenString)

	if err != nil {
		return 0, err
	}

	controller.UserLastActiveTime(claims.UserID)

	return claims.UserID, nil
}

// SigninRequired 必须是登录用户
func SigninRequired(ctx iris.Context) {

	user, err := getUser(ctx)
	if err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	ctx.Values().Set("user", user)
}

// EditorRequired 必须是网站编辑
func EditorRequired(ctx iris.Context) {

	user, err := getUser(ctx)
	if err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleEditor || user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		ctx.Values().Set("user", user)
	} else {
		common.Response(ctx, "没有权限", 1003)
	}
}

// AdminRequired 必须是管理员
func AdminRequired(ctx iris.Context) {

	user, err := getUser(ctx)
	if err != nil {
		common.Response(ctx, "未登录", e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		ctx.Values().Set("user", user)
	} else {
		common.Response(ctx, "没有权限", 1003)
	}
}

//Config全局变量太大了
//var jwtSecret = utils.ToBytes(initialize.Config.Server.JwtSecret))
