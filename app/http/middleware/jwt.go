package middleware

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/pkg/utils"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := utils.VerifyToken(c); err == nil {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
