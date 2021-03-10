package router

import (
	"app/infrastructure/middleware"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

type router struct {
	Engin *gin.Engine
}

func NewRouter(h *handler.Server, p http.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Auth())
	r.Use(middleware.Cors())

	r.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})
	return r
}

func (r *router) Run() {
	r.Engin.Run()
}
