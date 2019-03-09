package router

//go:generate qtc -dir=../template
import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"
	. "github.com/kataras/iris/middleware/recover"
	"hoper/client/controller"
	"hoper/client/controller/common/logging"
	"hoper/client/controller/hnsq"
	"hoper/client/controller/upload"
	"hoper/client/middleware"
	"hoper/client/router/other"
	"strings"
	"time"
)

func init() {
	//raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func IrisRouter() *iris.Application {
	app := iris.New()

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//关闭所有主机
		app.Shutdown(ctx)
	})
	app.StaticWeb("/api/static", "../static")
	app.Use(New())
	//other.Wrap(app)
	//api文档
	//other.Api(app)
	//https://rpm.newrelic.com/accounts/2269290/applications
	/*	config := newrelic.Config("hoper", "199e00247f278548fe92d6c81aeaadac0fc52b4b")
		m, err := newrelic.New(config)
		if err != nil {
			app.Logger().Fatal(err)
		}
		app.Use(m.ServeHTTP)*/

	/*	prometheus := prometheus.New("hoper")
		app.Use(prometheus.ServeHTTP)
		app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
			//错误代码处理程序不与其他路由共享相同的中间件，所以单独执行错误
			prometheus.ServeHTTP(ctx)
			ctx.Writef("Not Found")
		})
	*/
	/*middleware必须要写ctx.next(),且写在路由前，路由后的midddleware在请求之前的路由时不生效
	  iris.FromStd()将其他Handler转为iris的Handler
	*/
	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "../i18n/locale_en-US.ini",
			"zh-CN": "../i18n/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	logM := logMid()

	app.Use(logM)
	/*
		app.OnErrorCode(404 ,customLogger, func(ctx iris.Context) {
		   ctx.Writef("My Custom 404 error page ")
		})
	*/
	app.OnAnyErrorCode(logM, func(ctx iris.Context) {
		//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
		ctx.Values().Set("logger_message",
			"a dynamic message passed to the logs")
		ctx.Writef("My Custom error page")
	})
	//app.Logger().Printer.SetOutput(logging.F)

	UserRouter(app)

	WS(app)

	ArticleRouter(app)

	MomentRouter(app)
	//试验性
	other.GraphqlRouter(app)
	TPLRouter(app)
	other.Smart(app)
	//自己做还是第三方库刷新writer都没用
	//other.Sse(app)

	app.Post("/api/upload/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.Upload(ctx)
	})
	app.Post("/api/upload_multiple/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.UploadMultiple(ctx)
	})

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	app.Post("/api/comment/:classify", middleware.JWT, controller.AddComment)

	app.Get("/api/push", controller.Push)

	app.Post("/api/nsq", hnsq.Start)

	app.Get("/api/shutdown", middleware.JWT, func(c iris.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		app.Shutdown(ctx)
	})
	return app
}

func logMid() iris.Handler {
	var excludeExtensions = [...]string{
		".js",
		".css",
		".jpg",
		".png",
		".ico",
		".svg",
	}

	c := logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	}

	logFile := logging.F

	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}
	//我们不想使用记录器，一些静态请求等
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h := logger.New(c)
	return h
}
