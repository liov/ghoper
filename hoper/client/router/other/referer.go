package other

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Referer(app *iris.Application) {
	app.Get("/api/referer", func(ctx context.Context) /*或iris.Context，Go 1.9+也是如此*/ {
		// GetReferrer提取并返回指定的Referer标头中的信息
		//在https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy中或通过URL查询参数是referer。
		r := ctx.GetReferrer()
		switch r.Type {
		case context.ReferrerSearch:
			ctx.Writef("Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google: %s\n", r.GoogleType)
		case context.ReferrerSocial:
			ctx.Writef("Social %s\n", r.Label)
		case context.ReferrerIndirect:
			ctx.Writef("Indirect: %s\n", r.URL)
		}
	})
	//URL查询参数是referer
	// http://localhost:8080?referer=https://twitter.com/Xinterio/status/1023566830974251008
	// http://localhost:8080?referer=https://www.google.com/search?q=Top+6+golang+web+frameworks&oq=Top+6+golang+web+frameworks
}
