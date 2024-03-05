package server

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/healed/internal/router"
)

type Server struct {
	Port   string
	Engine *gin.Engine
	Router router.Router
}

// Constructor Function
func NewServer(port string, engine *gin.Engine, router router.Router) *Server {
	return &Server{
		Port:   port,
		Engine: engine,
		Router: router,
	}
}

// Start function
func (s *Server) Start() {

	s.Engine.GET("/health", func(ctx *gin.Context) {
		s.Router.GetHealth(ctx)
	})
	s.Engine.Run(s.Port)
}
