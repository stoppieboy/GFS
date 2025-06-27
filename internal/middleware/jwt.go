package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stoppieboy/gfs/internal/service"
	"go.uber.org/zap"
)

func JWTMiddleware(authService service.AuthService, logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		logger.Infof(authHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		logger.Infof("Token recieved: %s", tokenString)
		token, err := authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Next()
	}
}
