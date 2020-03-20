package v1

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/pkg/utils"
	"net/http"
)

func LoginIndex(c *gin.Context) {
	user := models.User{}
	token, _ := utils.GenToken(user)
	c.JSON(http.StatusOK, gin.H{"token": token, "data": user})
}

// 更新token
func Refresh(c *gin.Context) (string, error) {
	claims, _ := utils.VerifyToken(c)
	return utils.GenToken(claims.User)
}

func JwtAuth(c *gin.Context) {
	if _, err := utils.VerifyToken(c); err == nil {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 4001})
		c.Abort()
	}
}
