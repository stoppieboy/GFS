package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/config"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
	c   *config.Config
)

func uploadHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Warnf("File Upload error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			log.Errorf("Failed to create upload directory: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare upload directory"})
			return
		}
	}

	dst := filepath.Join(file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		log.Errorf("Failed to save file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	log.Infof("File uploaded: %s", dst)
	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "Path": dst})
}

func downloadHandler(ctx *gin.Context) {
	
}

func RegisterFileRoutes(r *gin.Engine, cfg *config.Config, logger *zap.SugaredLogger) {
	log = logger
	c = cfg
	router := r.Group("/file")
	router.GET("/upload", uploadHandler)
	router.GET("/download", downloadHandler)
}
