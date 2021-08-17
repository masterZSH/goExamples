package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{
						Name: "user",
						Fields: graphql.Fields{
							"id": &graphql.Field{
								Type: graphql.String,
							},
							"b": &graphql.Field{
								Type: graphql.String,
							},
						},
					}),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return nil, nil
					},
				},
			},
		}),
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
