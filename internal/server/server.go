package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/config"
	"github.com/stoppieboy/gfs/internal/router"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	log    *zap.SugaredLogger
}

func New(cfg *config.Config, logger *zap.SugaredLogger) *Server {
	s := &Server{
		config: cfg,
		log: logger,
		router: gin.Default(),
	}
	s.routes()
	return s
}

func (s Server) routes() {
	router.RegisterFileRoutes(s.router, s.config, s.log)
}

func (s Server) Start() {
	s.log.Infof("Starting server on port %s", s.config.Port)
	err := s.router.Run(":"+s.config.Port)
	if err != nil {
		s.log.Fatal("Server failed to start: %v", err)
	}
}