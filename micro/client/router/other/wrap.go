package other

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"net/http"
	"runtime/debug"
)

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

		router(w, r)
	})
}
