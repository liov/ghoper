package other

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"strings"
)

func init() {
	//raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
	//go http.ListenAndServe("localhost:8080", nil)
}

func Wrap(app *iris.Application) {
	//路由装饰
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		if strings.Contains(r.URL.Path, "debug") {
			http.DefaultServeMux.ServeHTTP(w, r)
			return
		}

		router(w, r)
	})
}

func PProf(app *iris.Application) {
	//无效，只有路由包装有效
	/*	app.Get("/debug/pprof/", func(c iris.Context) {
			http.DefaultServeMux.ServeHTTP(c.ResponseWriter(), c.Request())
		})
	*/
	//这个的底层实现就是上面，为啥无效
	//app.Get("/debug/pprof/", iris.FromStd(http.DefaultServeMux))
	pprofRouter := app.Party("/api/debug/pprof")
	{
		//Any方法写/不写/是有区别的，现在看来是必须有/，具体的http方法不需要，至少Get实测不需要
		//Any的方法中有个路径处理方法，返回的是路径数组，如果路径是""，返回的是nil，无法添加Handler
		//具体的http方法及时加了末尾/也会在处理中去掉
		// 这里之所以这么写，是因为pprof的坑
		//http.HandleFunc("/debug/pprof/", Index)
		pprofRouter.Get("/{action:string}", func(c iris.Context) {
			//切掉/api
			if c.Params().Get("action") == "index" {
				c.Request().URL.Path = c.Request().URL.Path[len("/api") : len(c.Request().URL.Path)-5]
			} else {
				c.Request().URL.Path = c.Request().URL.Path[len("/api"):]
			}
			c.Next()
		}, iris.FromStd(http.DefaultServeMux))
	}
}
