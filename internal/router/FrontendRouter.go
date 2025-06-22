package router

import (
	// "path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func serveHTML(ctx *gin.Context) {
	// var template string = filepath.Join("templates", "index.html")
	ctx.HTML(200, "index.html", gin.H{"documentName": "some Name"})
}

func RegisterFrontend(r *gin.Engine, log *zap.SugaredLogger) {
	r.GET("/somepath", serveHTML)
	log.Infof("Here in frontendrouter")
}
