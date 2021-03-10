package handler

import (
	"app/graph"
	"app/graph/directives"
	"app/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func NewGraphQLHandler(r *graph.Resolver) *handler.Server {
	c := generated.Config{
		Resolvers: r,
	}
	c.Directives.Authentication = directives.Authentication

	return handler.NewDefaultServer(generated.NewExecutableSchema(c))
}

func NewPlayGroundHandler() http.Handler {
	return playground.Handler("GraphQL", "/query")
}
