package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/web"
	"leeblog.com/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

func BlogDetail(c *gin.Context) {
	if id, err := strconv.Atoi(strings.TrimSuffix(c.Param("id.html"), ".html")); err != nil {
	} else {
		blog := web.NewBlog().GetBlogById(id)
		blog.View += 1
		web.NewBlog().UpdateBlog(blog)
		utils.View(c, "/web/blog/detail.html", gin.H{"blog": blog})
	}
}
func BlogAdd(c *gin.Context) {
	childCategory := web.NewCategory().GetChildCategory()
	utils.View(c, "/web/blog/add.html", gin.H{"childCategory": childCategory})
}
func BlogDoAdd(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.PostForm("category"))
	user, _ := c.Get("user")
	blog := models.Blog{
		Title:       c.PostForm("title"),
		Content:     c.PostForm("content"),
		CategoryId:  categoryId,
		Description: c.PostForm("description"),
		UserId:      int(user.(models.User).ID),
	}
	web.NewBlog().AddBlog(blog)
	c.Redirect(http.StatusFound, "/")
}
func BlogEdit(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		utils.View404(c)
	} else {
		blog := web.NewBlog().GetBlogById(id)
		user, _ := c.Get("user")
		if user.(models.User).ID != uint(blog.UserId) {
			utils.View404(c)
		}
		childCategory := web.NewCategory().GetChildCategory()
		utils.View(c, "/web/blog/edit.html", gin.H{"blog": blog, "childCategory": childCategory})
	}
}
func BlogUpdate(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		utils.View404(c)
	} else {
		blog := web.NewBlog().GetBlogById(id)
		if blog.ID <= 0 {
			utils.View404(c)
		} else {

			blog.Title = c.PostForm("title")
			blog.Content = c.PostForm("content")
			blog.ID = uint(id)
			if categoryId, err := strconv.Atoi(c.PostForm("category")); err == nil {
				blog.CategoryId = categoryId
			}
			blog.Description = c.PostForm("description")
			web.NewBlog().UpdateBlog(blog)
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/blog/detail/%d", blog.ID))
	}
}
