package route

import (
	"app/infrastructure/database"
	"app/infrastructure/gcp"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engin *gin.Engine
}

func newServer() *Server {
	return &Server{
		Engin: gin.Default(),
	}
}

func Run() {
	s := newServer()
	conn := database.InitDB()
	client := gcp.InitClient()
	s.InitRouter(conn, client)
	s.Engin.Run()
}
