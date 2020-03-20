package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func View(c *gin.Context, tpl string, data ...map[string]interface{}) {
	h := gin.H{}
	if len(data) > 0 {
		for _, v := range data {
			for k, val := range v {
				h[k] = val
			}
		}
	}
	if v, ok := c.Get("categorySlice"); ok {
		h["categorySlice"] = v
	}
	if v, ok := c.Get("user"); ok {
		h["user"] = v
	}
	c.HTML(http.StatusOK, tpl, h)
}
func View404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "/web/error/404.html", nil)
}
