package common

import (
	"fastService/controller/binding"
	"github.com/valyala/fasthttp"
)

type HopeCtx fasthttp.RequestCtx

func (c *HopeCtx) ShouldBindWith() {

}

func BindWithJson(c *fasthttp.RequestCtx, obj interface{}) error {
	return binding.JSON.Bind(&c.Request, obj)
}
