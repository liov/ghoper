package controller

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"hoper/initialize"
	"net/http"
	"time"
)

var (
	// AES仅支持16,24或32字节的密钥大小。
	//您需要准确提供该密钥字节大小，或者从您键入的内容中获取密钥。
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

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
	}, iris.CookieEncode(sc.Encode))
}

func GetCookie(c iris.Context, name string) {
	c.GetCookie(name, iris.CookieDecode(sc.Decode))
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
