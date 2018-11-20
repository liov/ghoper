package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func Index(c *fasthttp.RequestCtx) {
	//fmt.Fprintf(c, "Hi there! RequestURI is %q", c.RequestURI())
	req := string(c.QueryArgs().Peek("s"))
	ress, _ := json.Marshal(map[string]interface{}{
		"code": User{},
		"msg":  req,
		"data": "3",
	})

	c.SetBody(ress)
	c.SetStatusCode(fasthttp.StatusOK)
}
