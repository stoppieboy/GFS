package main

import (
	"github.com/stoppieboy/gfs/internal/config"
	"github.com/stoppieboy/gfs/internal/logger"
	"github.com/stoppieboy/gfs/internal/server"
)

func main() {
	// gin
	cfg := config.Load()
	logger := logger.New(cfg.Env)
	s := server.New(cfg, logger)
	s.Start()
}
