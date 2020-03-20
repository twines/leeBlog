package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProfileIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1})
}
