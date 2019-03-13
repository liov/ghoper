package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/initialize"
	"hoper/model"
	"net/http"
	"time"
)

func JWT(ctx iris.Context) {
	code := e.SUCCESS
	user, err := getUser(ctx)

	if err != nil && err.Error() == "未登录" {
		code = e.ErrorAuthCheckTokenFail
	} else if err != nil && err.Error() == "登录超时" {
		code = e.ErrorAuthCheckTokenTimeout
	}

	if code != e.SUCCESS {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    "del",
			Path:     "/",
			Domain:   "hoper.xyz",
			Expires:  time.Now().Add(-1),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
		})
		common.Response(ctx, err.Error(), code)
		return
	}
	ctx.Values().Set("userId", user.ID)
	ctx.Values().Set("user", *user)
	ctx.Next()
}

func Login(ctx iris.Context) {
	code := e.SUCCESS
	user, err := getUser(ctx)

	if err != nil && err.Error() == "未登录" {
		code = e.ErrorAuthCheckTokenFail
	} else if err != nil && err.Error() == "登录超时" {
		code = e.ErrorAuthCheckTokenTimeout
	}

	if code != e.SUCCESS {
		ctx.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    "del",
			Path:     "/",
			Domain:   "hoper.xyz",
			Expires:  time.Now().Add(-1),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
		})
		common.Response(ctx, err.Error(), code)
		return
	}
	ctx.Values().Set("userId", user.ID)
	ctx.Values().Set("user", *user)
	ctx.Next()
}

//中间件的两种方式
func GetUser() iris.Handler {
	return func(ctx iris.Context) {
		user, err := getUser(ctx)
		if err != nil {
			common.Response(ctx, "请重新登录", 401)
			return
		}
		ctx.Values().Set("user", user) //坑哦，这里存的是指针哦，虽然这个函数没用过
		ctx.Next()
	}
}

func getUser(ctx iris.Context) (*controller.User, error) {

	tokenString := ctx.GetCookie("token")
	if len(tokenString) == 0 {
		tokenString = ctx.GetHeader("Authorization")
	}
	if len(tokenString) == 0 {
		return nil, errors.New("未登录")
	}

	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected login method: %v", token.Header["alg"])
		}
		return []byte(initialize.Config.Server.TokenSecret), nil
	})

	if tokenErr != nil {
		return nil, errors.New("未登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		user, err := controller.UserFromRedis(userID)
		if err != nil {
			return nil, errors.New("登录超时")
		}
		return &user, nil
	}
	return nil, errors.New("未登录")
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

var jwtSecret = []byte(initialize.Config.Server.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hoper",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
