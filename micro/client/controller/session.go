package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"hoper/model/ov"
)

var Sess *sessions.Sessions

//废的，所有东西存在cookie里，随便存个什么都太长了
//var Gsess  = gss.NewCookieStore([]byte("the-big-and-secret-fash-key-here"))
//var Gsess *memstore.MemStore

func SessSet(ctx iris.Context) {
	s := Sess.Start(ctx)
	s.SetFlash("user", ov.User{Name: "贾一饼"})
	ctx.Writef("Message setted, is available for the next request")
}

func SessGet(ctx iris.Context) {
	s := Sess.Start(ctx)
	user := s.GetFlash("user").(User)
	ctx.Writef("Hello %s", user.Name)
}

func SessTest(ctx iris.Context) {
	s := Sess.Start(ctx)
	user := s.GetFlash("user").(User)
	ctx.Writef("Ok you are coming from /set ,the value of the name is %s", user.Name)
	ctx.Writef(", and again from the same context: %s", user.Name)
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

//gob: type not registered for interface: controller.User
/*func GsessSet(ctx iris.Context) {
	session, _ := Gsess.Get(ctx.Request(), "hopergsid")
	// Set some session values.
	session.Values["user"] = User{Name: "贾一饼"} //the value is too long
	session.Values[1] = 2
	// Save it before we write to the response/return from the handler.
	session.Save(ctx.Request(), ctx.ResponseWriter())
	ctx.Writef("Message setted, is available for the next request")
}

func GsessGet(ctx iris.Context) {
	session, _ := Gsess.Get(ctx.Request(), "hopergsid")
	// Set some session values.
	user := session.Values["user"].(*User)
	ctx.Writef("Hello %s %d", user.Name, session.Values[1])
}*/
