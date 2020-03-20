package web

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/services/web"
	"leeblog.com/pkg/utils"
	"math"
	"strconv"
)

func Index(c *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	id := 0
	if i, err := strconv.Atoi(c.Param("id")); err == nil {
		id = i
	}
	blogSlice, count := web.NewBlog().GetBlogList(id, page)
	goodBlogSlice := web.NewBlog().GetGoodBlog()

	paginate := utils.Paginate("/", int(math.Ceil(float64(count)/15)), page)
	utils.View(c, "/web/index/index.html", gin.H{"blogSlice": blogSlice, "goodBlogSlice": goodBlogSlice, "count": count, "navId": id, "paginate": paginate})
}
