package common

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Index(c *fasthttp.RequestCtx) {
	fmt.Fprintf(c, "Hi there! RequestURI is %q", c.RequestURI())
}
