package router

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
	"hoper/client/router/user"
	"strconv"
)

type UserResolver struct {
	user User
}

func (u *UserResolver) ID() graphql.ID { return graphql.ID(strconv.Itoa(u.user.ID)) }
func (u *UserResolver) Name() string   { return u.user.Name }
func (u *UserResolver) Sex() string    { return u.user.Sex }
func (u *UserResolver) Phone() string  { return u.user.Phone }

type User struct {
	ID    int
	Name  string
	Sex   string
	Phone string
}

type Resolver struct {
}

func (r *Resolver) GetUser(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	id, _ := strconv.Atoi(string(args.ID))

	user := User{ID: id, Name: "一二三", Sex: "男", Phone: "666"}

	s := UserResolver{
		user: user,
	}

	return &s, nil
}

func UserRouter(app *iris.Application) {

	s := `
                schema {
                      query: Query
                }
				type Query {
  					getUser(id: ID!): User
				}
                type User {
                        ID:ID!
						Name:String!
						Sex:String!
						Phone:String!
                }
        `

	schema := graphql.MustParseSchema(s, &Resolver{}, graphql.UseStringDescriptions())
	app.Use(iris.FromStd(&relay.Handler{Schema: schema}))

	userRouter := app.Party("/api/user")
	{
		userRouter.Get("/active/{id:uint64}/{secret:string}", controller.ActiveAccount)
		userRouter.Post("/signup", controller.Signup)
		userRouter.Post("/login", controller.Login)
		userRouter.Get("/logout", middleware.JWT, controller.Logout)
		userRouter.Post("/active", controller.ActiveSendMail)
		userRouter.Get("/Get", middleware.JWT, controller.SignInFlag)
		userRouter.Post("/rpc/login", user.Login)
		userRouter.Get("/rpc/logout", user.Logout)
		userRouter.Post("/rpc/signup", user.Signup)
	}
}
