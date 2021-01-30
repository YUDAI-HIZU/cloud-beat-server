package route

import (
	"app/graph"
	"app/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Server) InitRouter(conn *gorm.DB) {
	s.Engin.Use(gin.Logger())
	s.Engin.Use(gin.Recovery())

	h := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				DB: conn,
			},
		}),
	)

	p := playground.Handler("GraphQL", "/query")

	s.Engin.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	s.Engin.GET("/", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
}
