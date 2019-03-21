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
	app.Get("/debug/pprof/", iris.FromStd(http.DefaultServeMux))

}
