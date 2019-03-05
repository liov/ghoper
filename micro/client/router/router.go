package router

//go:generate qtc -dir=../template
import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	. "github.com/kataras/iris/middleware/recover"
	"hoper/client/controller"
	"hoper/client/controller/common/logging"
	"hoper/client/controller/upload"
	"hoper/initialize"
	"hoper/model"
	"time"

	"hoper/client/controller/hnsq"
	"hoper/client/controller/hwebsocket"

	"hoper/client/middleware"
	"net/http"
	"runtime/debug"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
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

	app.Use(New())
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

	app.Logger().Printer.SetOutput(logging.F)

	WS(app)

	TPLRouter(app)

	ArticleRouter(app)

	MomentRouter(app)

	GraphqlRouter(app)
	app.Post("/api/upload/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.Upload(ctx)
	})
	app.Post("/api/upload_multiple/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.UploadMultiple(ctx)
	})

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	UserRouter(app)

	app.Post("/api/comment/:classify", middleware.JWT, controller.AddComment)

	//app.Get("/api/push",controller.Push)

	app.Get("/api/chat/getChat", hwebsocket.GetChat)

	app.Post("/api/nsq", hnsq.Start)

	app.Get("/api/init", func(c iris.Context) {
		initialize.DB.DropTableIfExists(&model.User{},
			&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.School{}, &model.Article{},
			&model.Career{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
			&model.ArticleComment{},
			&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
			&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
			&model.MomentHistory{}, "article_category", "article_collection", "article_comment", "article_history",
			"article_history_category", "article_history_tag", "article_like", "article_tag",
			"diary_book_category", "diary_book_collection", "diary_book_comment", "diary_book_history",
			"diary_book_history_category", "diary_book_history_tag", "diary_book_like", "diary_book_tag",
			"diary_category", "diary_collection", "diary_comment", "diary_history",
			"diary_history_category", "diary_history_tag", "diary_like", "diary_tag",
			"moment_category", "moment_collection", "moment_comment", "moment_history",
			"moment_history_category", "moment_history_tag", "moment_like", "moment_tag", "user_collection")
		initialize.DB.CreateTable(&model.User{},
			&model.Tag{}, &model.Diary{}, &model.DiaryBook{}, &model.School{}, &model.Article{},
			&model.Career{}, &model.Category{}, &model.Collection{}, &model.Like{}, &model.CrawlerArticle{},
			&model.ArticleComment{},
			&model.DiaryBookComment{}, &model.MomentComment{}, &model.FileUploadInfo{}, &model.DiaryComment{},
			&model.Moment{}, &model.Mood{}, &model.DiaryBookHistory{}, &model.DiaryHistory{}, &model.ArticleHistory{},
			&model.MomentHistory{})
	})

	return app
}
