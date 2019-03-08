package other

import (
	"encoding/xml"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Data(app *iris.Application) {
	dataR := app.Party("/api/data")
	dataR.Post("/decode", func(ctx iris.Context) {
		// 参考 /http_request/read-json/main.go
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s is %d years old and comes from %s!", user.Name, user.Phone, user.Sex, user.ID)
	})
	// Write
	dataR.Get("/encode", func(ctx iris.Context) {
		peter := User{
			Name:  "John",
			Phone: "186XXXX2064",
			Sex:   "男",
			ID:    25,
		}
		//手动设置内容类型: ctx.ContentType("application/javascript")
		ctx.JSON(peter)
	})
	//其他内容类型
	dataR.Get("/binary", func(ctx iris.Context) {
		//当您想要强制下载原始字节内容时有用下载文件
		ctx.Binary([]byte("Some binary data here."))
	})
	dataR.Get("/text", func(ctx iris.Context) {
		ctx.Text("Plain text here")
	})
	dataR.Get("/json", func(ctx iris.Context) {
		ctx.JSON(map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
	})
	dataR.Get("/jsonp", func(ctx iris.Context) {
		ctx.JSONP(map[string]string{"hello": "jsonp"}, context.JSONP{Callback: "callbackName"})
	})
	dataR.Get("/xml", func(ctx iris.Context) {
		ctx.XML(ExampleXML{One: "hello", Two: "xml"}) // or iris.Map{"One":"hello"...}
	})
	dataR.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
	})
	// http://localhost:8080/decode
	// http://localhost:8080/encode
	// http://localhost:8080/binary
	// http://localhost:8080/text
	// http://localhost:8080/json
	// http://localhost:8080/jsonp
	// http://localhost:8080/xml
	// http://localhost:8080/markdown

	//`iris.WithOptimizations`是一个可选的配置器，
	//如果传递给`Run`那么它将确保应用程序，尽快响应客户端。
}

// ExampleXML只是一个要查看的测试结构，表示xml内容类型
type ExampleXML struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}
