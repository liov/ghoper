package router

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/kataras/iris"
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
		graphqlRouter.Post("/", iris.FromStd(&relay.Handler{Schema: schema}))
	}
}
