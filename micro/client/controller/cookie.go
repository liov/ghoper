package controller

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"hoper/initialize"
	"net/http"
	"time"
)

var secureCookie *securecookie.SecureCookie

func SetCookie(c iris.Context, key, value string) {
	c.SetCookie(&http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		Domain:   "hoper.xyz",
		Expires:  time.Now().Add(time.Duration(initialize.Config.Server.TokenMaxAge) * time.Second),
		MaxAge:   int(time.Duration(initialize.Config.Server.TokenMaxAge) * time.Second),
		Secure:   false,
		HttpOnly: true,
	}, iris.CookieEncode(secureCookie.Encode))
}

func GetCookie(c iris.Context, name string) {
	c.GetCookie(name, iris.CookieDecode(secureCookie.Decode))
}

func DeleteCookie(c iris.Context, key string) {
	c.SetCookie(&http.Cookie{
		Name:     key,
		Value:    "del",
		Path:     "/",
		Domain:   "hoper.xyz",
		Expires:  time.Now().Add(-1),
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})
}
