package router

import (
	"app/graph"
	"app/graph/directives"
	"app/graph/generated"
	"app/infrastructure/database"
	"app/infrastructure/middleware"
	"app/infrastructure/storage"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type router struct {
	Engin *gin.Engine
}

func NewRouter() *router {
	r := &router{
		Engin: gin.Default(),
	}

	db := database.NewDatabase()

	storage := storage.NewStorageClient()

	r.Engin.Use(gin.Logger())
	r.Engin.Use(gin.Recovery())
	r.Engin.Use(middleware.Auth())
	r.Engin.Use(middleware.Cors())
	c := generated.Config{
		Resolvers: &graph.Resolver{
			DB:      db.Client,
			Storage: storage.Client,
		},
	}

	c.Directives.Authentication = directives.Authentication

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	p := playground.Handler("GraphQL", "/query")

	r.Engin.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Engin.GET("/", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
	return r
}

func (r *router) Run() {
	r.Engin.Run()
}
