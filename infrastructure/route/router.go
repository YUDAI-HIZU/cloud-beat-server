package route

import (
	"app/graph"
	"app/graph/directives"
	"app/graph/generated"
	"app/infrastructure/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Server) InitRouter(conn *gorm.DB) {
	s.Engin.Use(gin.Logger())
	s.Engin.Use(gin.Recovery())
	s.Engin.Use(middleware.Auth())
	s.Engin.Use(middleware.Cors())

	c := generated.Config{
		Resolvers: &graph.Resolver{
			DB: conn,
		},
	}

	c.Directives.Authentication = directives.Authentication

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	p := playground.Handler("GraphQL", "/query")

	s.Engin.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	s.Engin.GET("/", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
}
