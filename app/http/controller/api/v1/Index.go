package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags 首页
// @Summary 首页
// @Description # aaaaa
// @Produce  json
// @Param id path int true "ID"
// @Param name query string false "ID"
// @Param img formData file false "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/ [get]
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"a": "a"})
}

// @Tags 首页
// @Summary 详情
// @Description id 文章id|name文章名称
// @Produce  json
// @Param id path int true "ID"
// @Param name query string false "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/index/detail [get]
func IndexDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"a": "a"})
}

// IndexPost ndexpost
// @Tags 首页
// @Summary 详情
// @Description id 文章id|name文章名称
// @Produce  json
// @Param id path int true "ID"
// @Param img formData file false "img"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/index/post [post]
// @deprecated true
func IndexPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"a": "a"})
}
