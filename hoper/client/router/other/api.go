package other

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

func Api(app *iris.Application) {
	yaag.Init(&yaag.Config{
		On:       true, //是否开启自动生成API文档功能
		DocTitle: "Iris",
		DocPath:  "../../static/api/apidoc.html", //生成API文档名称存放路径
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	//注册中间件
	app.Use(irisyaag.New())
}
