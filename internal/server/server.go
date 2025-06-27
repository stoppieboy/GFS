package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/config"
	"github.com/stoppieboy/gfs/internal/middleware"
	"github.com/stoppieboy/gfs/internal/router"
	"github.com/stoppieboy/gfs/internal/service"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	log    *zap.SugaredLogger
}

var(
	uploadDir = "uploads"
	secret = "secret"
)

func New(cfg *config.Config, logger *zap.SugaredLogger) *Server {
	s := &Server{
		config: cfg,
		log: logger,
		router: gin.Default(),
	}
	//load html templates
	s.router.LoadHTMLGlob("templates/*")
	//load static content
	s.router.Static("/static", "./static")
	s.routes()
	return s
}

func (s Server) routes() {
	fileService, err := service.NewFileService(uploadDir)
	if err != nil {
		s.log.Errorf("Failed to create file service: %v", err)
		return
	}
	authService := service.NewAuthService(secret)
	router.RegisterAuthRoutes(s.router, s.log, authService)

	authMiddleware := middleware.JWTMiddleware(authService, s.log)
	authGroup := s.router.Group("/file")
	authGroup.Use(authMiddleware)
	router.RegisterFileRoutes(authGroup, s.config, s.log, fileService)
	router.RegisterFrontend(s.router, s.log)
}

func (s Server) Start() {
	s.log.Infof("Starting server on port %s", s.config.Port)
	err := s.router.Run(":"+s.config.Port)
	if err != nil {
		s.log.Fatal("Server failed to start: %v", err)
	}
}