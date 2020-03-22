package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/admin"
	"leeblog.com/pkg/utils"
	"net/http"
	"strconv"
)

func NewsAdd(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	keyword := c.PostForm("keyword")
	description := c.PostForm("description")
	news := models.News{}
	news.Title = title
	news.Content = content
	news.Description = description
	news.Keyword = keyword
	if admin.NewNewsService().AddNews(news) {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "添加成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "添加失败，请稍后重试"})
}
func NewsDetail(c *gin.Context) {
	newsId := 0
	if id, err := strconv.Atoi(c.Param("newsId")); err == nil {
		newsId = id
	}
	if newsId <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "newsId必须大于0"})
		return
	}
	news := admin.NewNewsService().GetNewsById(newsId)
	if news.ID < 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "新闻公告不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": news})
}
func NewsDelete(c *gin.Context) {
	newsId := 0
	if id, err := strconv.Atoi(c.Param("newsId")); err == nil {
		newsId = id
	}
	if newsId <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "newsId必须大于0"})
		return
	}
	news := admin.NewNewsService().GetNewsById(newsId)
	if news.ID < 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "新闻公告不存在"})
		return
	}
	if admin.NewNewsService().DeleteNews(newsId) {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "删除失败"})
	}
}
func NewsUpdate(c *gin.Context) {

}
func NewsList(c *gin.Context) {
	page, limit := 1, 2
	if p, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	newsSlice, count := admin.NewNewsService().GetNewsList(page, limit)
	totalPage := utils.Page(count, limit)
	mp := map[string]interface{}{}
	mp["totalCount"] = count
	mp["limit"] = limit
	mp["totalPage"] = totalPage
	mp["data"] = newsSlice
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}
