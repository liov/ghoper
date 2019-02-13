package common

import (
	"github.com/valyala/fasthttp"
	"fastService/controller/binding"
)

type HopeCtx fasthttp.RequestCtx

func (c *HopeCtx) ShouldBindWith() {

}

func BindWithJson(c *fasthttp.RequestCtx, obj interface{}) error {
	return binding.JSON.Bind(&c.Request, obj)
}
