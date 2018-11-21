package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"service/controller/common"
)

func Index(c *fasthttp.RequestCtx) {
	//fmt.Fprintf(c, "Hi there! RequestURI is %q", c.RequestURI())
	req := string(c.QueryArgs().Peek("s"))
	ress, _ := json.Marshal(common.H{
		"code": User{},
		"msg":  req,
		"data": "中文",
	})

	c.SetBody(ress)
	c.SetStatusCode(fasthttp.StatusOK)
}
