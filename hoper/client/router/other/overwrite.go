package other

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/sessions"
	"strings"
	"sync"
)

func OverWrite(app *iris.Application) {

	app.ContextPool.Attach(func() context.Context {
		return &MyContext{
			// Optional Part 3:
			Context: context.NewContext(app),
		}
	})

	app.Post("/api/set", Handler(func(ctx *Context) {
		nameFieldValue := ctx.FormValue("name")
		ctx.Session().Set("name", nameFieldValue)
		ctx.Writef("set session = " + nameFieldValue)
	}))

	app.Get("/api/cs/{num:uint64 min(10) else 400}", func(ctx context.Context) {
		num := ctx.Params().GetUint64Default("num", 0)
		ctx.Writef("num is: %d\n", num)
	})
	myCustomRouter := new(customRouter)
	app.BuildRouter(app.ContextPool, myCustomRouter, app.APIBuilder, true)
}

type MyContext struct {
	// Optional Part 1: embed (optional but required if you don't want to override all context's methods)
	context.Context // it's the context/context.go#context struct but you don't need to know it.
}

var _ context.Context = &MyContext{} // optionally: validate on compile-time if MyContext implements context.Context.

// The only one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the handlers via this "*MyContext".
func (ctx *MyContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

func (ctx *MyContext) Next() {
	context.Next(ctx)
}

func (ctx *MyContext) HTML(htmlContents string) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from MyContext")

	ctx.ContentType("text/html")
	return ctx.WriteString(htmlContents)
}

type Owner struct {
	// define here the fields that are global
	// and shared to all clients.
	sessionsManager *sessions.Sessions
}

// this package-level variable "application" will be used inside context to communicate with our global Application.
var owner = &Owner{
	sessionsManager: sessions.New(sessions.Config{Cookie: "mysessioncookie"}),
}

// Context is our custom context.
// Let's implement a context which will give us access
// to the client's Session with a trivial `ctx.Session()` call.
type Context struct {
	iris.Context
	session *sessions.Session
}

// Session returns the current client's session.
func (ctx *Context) Session() *sessions.Session {
	// this help us if we call `Session()` multiple times in the same handler
	if ctx.session == nil {
		// start a new session if not created before.
		ctx.session = owner.sessionsManager.Start(ctx.Context)
	}

	return ctx.session
}

// Bold will send a bold text to the client.
func (ctx *Context) Bold(text string) {
	ctx.HTML("<b>" + text + "</b>")
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	ctx.session = nil      // reset the session
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

// Handler will convert our handler of func(*Context) to an iris Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

/* A Router should contain all three of the following methods:
   - HandleRequest should handle the request based on the Context.
	  HandleRequest(ctx context.Context)
   - Build should builds the handler, it's being called on router's BuildRouter.
	  Build(provider router.RoutesProvider) error
   - RouteExists reports whether a particular route exists.
      RouteExists(ctx context.Context, method, path string) bool
For a more detailed, complete and useful example
you can take a look at the iris' router itself which is located at:
https://github.com/kataras/iris/tree/master/core/router/handler.go
which completes this exact interface, the `router#RequestHandler`.
*/
type customRouter struct {
	// a copy of routes (safer because you will not be able to alter a route on serve-time without a `app.RefreshRouter` call):
	// []router.Route
	// or just expect the whole routes provider:
	provider router.RoutesProvider
}

// HandleRequest a silly example which finds routes based only on the first part of the requested path
// which must be a static one as well, the rest goes to fill the parameters.
func (r *customRouter) HandleRequest(ctx context.Context) {
	path := ctx.Path()
	ctx.Application().Logger().Infof("Requested resource path: %s", path)

	parts := strings.Split(path, "/")[1:]
	staticPath := "/" + parts[0]
	for _, route := range r.provider.GetRoutes() {
		if strings.HasPrefix(route.Path, staticPath) && route.Method == ctx.Method() {
			paramParts := parts[1:]
			for _, paramValue := range paramParts {
				for _, p := range route.Tmpl().Params {
					ctx.Params().Set(p.Name, paramValue)
				}
			}

			ctx.SetCurrentRouteName(route.Name)
			ctx.Do(route.Handlers)
			return
		}
	}

	// if nothing found...
	ctx.StatusCode(iris.StatusNotFound)
}

func (r *customRouter) Build(provider router.RoutesProvider) error {
	for _, route := range provider.GetRoutes() {
		// do any necessary validation or conversations based on your custom logic here
		// but always run the "BuildHandlers" for each registered route.
		route.BuildHandlers()
		// [...] r.routes = append(r.routes, *route)
	}

	r.provider = provider
	return nil
}

func (r *customRouter) RouteExists(ctx context.Context, method, path string) bool {
	// [...]
	return false
}
