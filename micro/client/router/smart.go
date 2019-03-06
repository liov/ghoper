package router

import (
	"fmt"
	"github.com/jmespath/go-jmespath"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/hero"
	"github.com/opentracing/opentracing-go/log"
	"hoper/initialize"
)

func Smart(app *iris.Application) {
	app.PartyFunc("/api/users", registerUsersRoutes)
}

func registerUsersRoutes(usersRouter iris.Party) {
	// GET: /users
	usersRouter.Get("/", getAllUsersHandler)
	usersRouter.PartyFunc("/{id:int}", registerUserRoutes)
}
func getAllUsersHandler(ctx iris.Context) {
	var users []User
	initialize.DB.Find(&users)
	err := sendJSON(ctx, users)
	if err != nil {
		log.Error(err)
	}
}

func registerUserRoutes(userRouter iris.Party) {
	//为此子路由器创建一个新的依赖注入管理器
	userDeps := hero.New()
	//你也可以使用global/package-level的hero.Register(userDependency)，正如我们在其他例子中已经学到的那样。
	userDeps.Register(userDependency)
	// GET: /users/{name:string}
	userRouter.Get("/", userDeps.Handler(getUserHandler))
	// GET: /users/{name:string}/age
	userRouter.Get("/name", userDeps.Handler(getUserNameHandler))
}

var userDependency = func(ctx iris.Context) *User {
	var user User

	id := ctx.Params().GetIntDefault("id", 0)

	initialize.DB.Where("id=?", id).Find(&user)
	return &user

	fail(ctx, iris.StatusNotFound, "user with name '%s' not found", id)
	return nil
}

func getUserHandler(ctx iris.Context, u *User) {
	if u == nil {
		return
	}
	sendJSON(ctx, u)
}

func getUserNameHandler(ctx iris.Context, u *User) {
	if u == nil {
		return
	}
	ctx.Writef("%d", u.Name)
}

type httpError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (h httpError) Error() string {
	return fmt.Sprintf("Status Code: %d\nReason: %s", h.Code, h.Reason)
}

func fail(ctx iris.Context, statusCode int, format string, a ...interface{}) {
	err := httpError{
		Code:   statusCode,
		Reason: fmt.Sprintf(format, a...),
	}
	//记录所有> = 500内部错误。
	if statusCode >= 500 {
		ctx.Application().Logger().Error(err)
	}
	ctx.StatusCode(statusCode)
	ctx.JSON(err)
	//没有下一个处理程序将运行。
	ctx.StopExecution()
}

func sendJSON(ctx iris.Context, resp interface{}) (err error) {
	indent := ctx.URLParamDefault("indent", "  ")
	// i.e [?Name == 'John Doe'].Age # to output the [age] of a user which his name is "John Doe".
	if query := ctx.URLParam("query"); query != "" && query != "[]" {
		//只能搜字符串?
		//http://hoper.xyz/api/users?query=[?ID%20!=%20`1`]有结果，卧槽，什么操作？？
		resp, err = jmespath.Search(query, resp)
		if err != nil {
			return
		}
	}
	_, err = ctx.JSON(resp, context.JSON{Indent: indent, UnescapeHTML: true})
	return err
}
