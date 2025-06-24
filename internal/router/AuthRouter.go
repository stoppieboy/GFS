package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/service"
	"go.uber.org/zap"
)

func RegisterAuthRoutes(r *gin.Engine, log *zap.SugaredLogger, authService service.AuthService) {
	r.POST("/login", func(ctx *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := ctx.ShouldBindJSON(&credentials); err != nil {
			log.Warnf("Invalid login payload: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		token, err := authService.Login(credentials.Username, credentials.Password)
		if err != nil {
			log.Warnf("Unauthorized: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "auth.html", gin.H{})
	})
}