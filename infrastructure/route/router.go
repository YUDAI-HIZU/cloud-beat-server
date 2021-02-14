package route

import (
	"app/config"
	"app/graph"
	"app/graph/directives"
	"app/graph/generated"
	"app/infrastructure/middleware"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/option"
)

func (s *Server) InitRouter(conn *gorm.DB, client *storage.Client) {
	s.Engin.Use(gin.Logger())
	s.Engin.Use(gin.Recovery())
	s.Engin.Use(middleware.Auth())
	s.Engin.Use(middleware.Cors())

	c := generated.Config{
		Resolvers: &graph.Resolver{
			DB:     conn,
			Client: client,
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
	s.Engin.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err)
		}
		imageFile, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err)
		}

		ctx := c.Request.Context()
		client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(config.GCSAccount)))
		if err != nil {
			fmt.Println(err)
		}

		writer := client.Bucket("cloud-beat-bucket").Object("production/user/").NewWriter(ctx)
		writer.ContentType = "image/png"
		ok, err := io.Copy(writer, imageFile)
		fmt.Println(ok, "=======okko========")
		if err != nil {
			fmt.Println(err)
		}
		err = writer.Close()
		if err != nil {
			fmt.Println(err)
		}
	})
}
