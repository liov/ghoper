package other

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kataras/iris"
	"hoper/client/router/user"
	"hoper/protobuf"
	"strconv"
)

func GraphqlRouter(app *iris.Application) {

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

	graphqlRouter := app.Party("/api/graphql")
	{
		graphqlRouter.Post("", iris.FromStd(&relay.Handler{Schema: schema}))
	}
}

type UserResolver struct {
	user protobuf.User
}

func (u *UserResolver) ID() graphql.ID { return graphql.ID(strconv.Itoa(int(u.user.ID))) }
func (u *UserResolver) Name() string   { return u.user.Name }
func (u *UserResolver) Sex() string    { return u.user.Sex }
func (u *UserResolver) Phone() string  { return u.user.Phone }

type Resolver struct {
}

func (r *Resolver) GetUser(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	id, _ := strconv.Atoi(string(args.ID))
	getReq := protobuf.GetReq{ID: uint64(id)}
	user, err := user.Service.GetUser(context.TODO(), &getReq)
	if err != nil {

		return nil, err
	}

	s := UserResolver{
		user: *user,
	}

	return &s, nil
}
