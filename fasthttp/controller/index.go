package controller

import (
	"github.com/valyala/fasthttp"
	"fastService/controller/common"
)

func Index(c *fasthttp.RequestCtx) {
	//fmt.Fprintf(c, "Hi there! RequestURI is %q", c.RequestURI())
	req := c.QueryArgs().Peek("s")
	reqs := "不是0"
	if len(req) != 0 {
		req := c.QueryArgs().Peek("s")[0] - 48
		if req == 0 {
			reqs = "是0"
		}
	}
	ress, _ := jsons.MarshalToString(common.H{
		"code": User{},
		"msg":  reqs,
		"data": "中文",
	})

	c.SetBodyString(ress)
	c.SetStatusCode(fasthttp.StatusOK)

}
