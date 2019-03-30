package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var Sess *sessions.Sessions

func SessSet(ctx iris.Context) {
	s := Sess.Start(ctx)
	s.SetFlash("user", User{Name: "贾一饼"})
	ctx.Writef("Message setted, is available for the next request")
}

func SessGet(ctx iris.Context) {
	s := Sess.Start(ctx)
	name := s.GetFlash("user").(User)
	ctx.Writef("Hello %s", name.Name)
}

func SessTest(ctx iris.Context) {
	s := Sess.Start(ctx)
	name := s.GetFlashString("user")
	if name == "" {
		ctx.Writef("Empty name!!")
		return
	}

	ctx.Writef("Ok you are coming from /set ,the value of the name is %s", name)
	ctx.Writef(", and again from the same context: %s", name)
}

func SessDelete(ctx iris.Context) {
	// delete a specific key
	Sess.Start(ctx).Delete("user")
}

func SessClear(ctx iris.Context) {
	// removes all entries
	Sess.Start(ctx).Clear()
}

func SessDestroy(ctx iris.Context) {
	//destroy, removes the entire Session data and cookie
	Sess.Destroy(ctx)
}

func SessUpdate(ctx iris.Context) {
	// updates resets the expiration based on the Session's `Expires` field.
	if err := Sess.ShiftExpiration(ctx); err != nil {
		if sessions.ErrNotFound.Equal(err) {
			ctx.StatusCode(iris.StatusNotFound)
		} else if sessions.ErrNotImplemented.Equal(err) {
			ctx.StatusCode(iris.StatusNotImplemented)
		} else {
			ctx.StatusCode(iris.StatusNotModified)
		}

		ctx.Writef("%v", err)
		ctx.Application().Logger().Error(err)
	}
}
