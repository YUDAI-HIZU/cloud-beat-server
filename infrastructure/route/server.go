package route

import (
	"app/infrastructure/database"

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
	s.InitRouter(conn)
	s.Engin.Run()
}
