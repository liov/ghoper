package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin/json"
	"github.com/valyala/fasthttp"
	"service/controller"
	"service/controller/common"
	"service/controller/common/e"
	"service/initialize"
	"service/model"
)

func JWT(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		code := e.SUCCESS
		user, err := getUser(c)
		if err != nil && err.Error() == "没有token"{
			code = e.InvalidParams
		} else if err != nil && err.Error() == "未登录" {
			code = e.ErrorAuthCheckTokenFail
		} else if err != nil && err.Error() == "登录超时" {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			c.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
			res,_:= json.Marshal(map[string]interface{}{
				"code": code,
				"data": e.GetMsg(code),})
			c.SetBody(res)
			return
		}
		c.SetUserValue("user", user)
		h(c)
	}
}

func GetUser(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		user,_ := getUser(c)
		c.SetUserValue("user", user)
		h(c)
	}
}

func getUser(c *fasthttp.RequestCtx) (controller.User, error) {
	var user controller.User
	tokenString := c.Request.Header.Cookie("token")
	if len(tokenString)==0 {
		tokenString = c.Request.Header.Peek("Authorization")
	}
	if len(tokenString)==0 {
		return user, errors.New("没有token")
	}

	token, tokenErr := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initialize.ServerSettings.TokenSecret), nil
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
func SetContextUser(c *fasthttp.RequestCtx) {
	var user controller.User
	var err error
	if user, err = getUser(c); err != nil {
		c.SetUserValue("user", nil)
		return
	}
	c.SetUserValue("user", user)
}

// SigninRequired 必须是登录用户
func SigninRequired(c *fasthttp.RequestCtx) {

	var user controller.User
	var err error
	if user, err = getUser(c); err != nil {
		common.Response(c, "未登录", e.LoginTimeout,)
		return
	}
	c.SetUserValue("user", user)
}

// EditorRequired 必须是网站编辑
func EditorRequired(c *fasthttp.RequestCtx) {

	var user controller.User
	var err error
	if user, err = getUser(c); err != nil {
		common.Response(c,  "未登录",e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleEditor || user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		c.SetUserValue("user", user)
	} else {
		common.Response(c,  "没有权限",1003)
	}
}

// AdminRequired 必须是管理员
func AdminRequired(c *fasthttp.RequestCtx) {

	var user controller.User
	var err error
	if user, err = getUser(c); err != nil {
		common.Response(c, "未登录", e.LoginTimeout)
		return
	}
	if user.Role == model.UserRoleAdmin || user.Role == model.UserRoleCrawler || user.Role == model.UserRoleSuperAdmin {
		c.SetUserValue("user", user)
	} else {
		common.Response(c,  "没有权限",1003)
	}
}
