package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/config"
	"github.com/stoppieboy/gfs/internal/service"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
	c   *config.Config
	s   service.FileService
)

func uploadHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Warnf("File Upload error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	path, err := s.Save(file)
	if err != nil {
		log.Errorf("Failed to save file: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	log.Infof("File uploaded: %s", path)
	// ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "Path": path})
	filepath, err := s.Get(file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
		return
	}
	log.Infof("here: %s", filepath)
	ctx.File(filepath)
}

func downloadHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	filepath, err := s.Get(filename)
	if err != nil {
		log.Errorf("File not found: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	ctx.File(filepath)
}

func deleteHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	err := s.Delete(filename)
	if err != nil {
		log.Warnf("Failed to delete the file: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found or failed to delete"})
		return
	}

	log.Infof("File deleted: %s", filename)
	ctx.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

func RegisterFileRoutes(router *gin.RouterGroup, cfg *config.Config, logger *zap.SugaredLogger, fileService service.FileService) {
	log = logger
	c = cfg
	s = fileService
	router.POST("/", uploadHandler)
	router.GET("/:filename", downloadHandler)
	router.DELETE("/:filename", deleteHandler)
}
