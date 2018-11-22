package common

import (
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"service/controller/common/e"
	"time"
)

type H map[string]interface{}

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func Response(c *fasthttp.RequestCtx, res ...interface{}) {

	var msg string
	var code int
	var data interface{}
	if len(res) == 1 {
		code = e.SUCCESS
		if msgTmp, ok := res[0].(string); ok {
			msg = msgTmp
			data = nil
		} else {
			data = res[0]
			code = e.SUCCESS
		}
	} else if len(res) == 2 {
		if msgTmp, ok := res[0].(string); ok {
			msg = msgTmp
			code = res[1].(int)
			data = nil
		} else {
			msg = res[1].(string)
			code = e.SUCCESS
			data = res[0]

		}
	} else {
		code = res[2].(int)
		msg = res[1].(string)
		data = res[0]
	}

	ress, _ := Json.MarshalToString(H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	c.SetBodyString(ress)
	c.SetStatusCode(fasthttp.StatusOK)
}

func Res(c *fasthttp.RequestCtx, h H) {
	res, _ := Json.MarshalToString(h)
	c.SetBodyString(res)
	c.SetStatusCode(fasthttp.StatusOK)
}

func SetCookie(key, value string, maxAge time.Duration, path, domain string, secure, httpOnly bool) *fasthttp.Cookie {
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey(key)
	cookie.SetHTTPOnly(httpOnly)
	cookie.SetDomain(domain)
	cookie.SetExpire(time.Now().Add(maxAge))
	return cookie
}
