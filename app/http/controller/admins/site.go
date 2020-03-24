package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/admin"
	"net/http"
	"strconv"
)

func SiteConfig(c *gin.Context) {
	siteName := c.PostForm("siteName")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	number := c.PostForm("number")
	status := c.PostForm("status")
	logo := c.PostForm("logo")
	site := models.Site{}
	site.SiteName = siteName
	site.Keywords = keywords
	site.Description = description
	site.Number = number
	if s, err := strconv.Atoi(status); err == nil {
		site.Status = uint8(s)
	}
	site.Logo = logo
	admin.NewSiteService().AddSiteConfig(site)
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
}

func SiteDetail(c *gin.Context) {
	site := admin.NewSiteService().GetSiteConfig()
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": site})
}
