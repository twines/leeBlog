package web

import (
	"leeblog.com/app/models"
)

type Blog struct {
}

func NewBlog() *Blog {
	return &Blog{}
}
func (blog *Blog) GetBlogList(categoryId int, page int) (blogSlice []models.Blog, count int) {
	if categoryId > 0 {
		models.DB().Joins("left join categories on blogs.category_id=categories.id").Order("id desc").Where("categories.parent_id=?", categoryId).Find(&Blog{}).Count(&count)
		models.DB().Joins("left join categories on blogs.category_id=categories.id").Order("id desc").Where("categories.parent_id=?", categoryId).Offset((page - 1) * 15).Limit(15).Find(&blogSlice)
	} else {
		models.DB().Find(&Blog{}).Count(&count)
		models.DB().Order("id desc").Offset((page - 1) * 15).Limit(15).Find(&blogSlice).Count(&count)
	}
	return
}

func (blog *Blog) GetGoodBlog() (blogSlice []models.Blog) {
	models.DB().Order("view desc,id desc").Limit(6).Find(&blogSlice)
	return
}
func (blog *Blog) GetBlogById(id int) (b models.Blog) {
	models.DB().First(&b, id)
	return
}
func (blog *Blog) AddBlog(b models.Blog) (id uint) {
	models.DB().Create(&b)
	return b.ID
}

func (blog *Blog) UpdateBlog(b models.Blog) (id int64) {
	d := models.DB().Save(&b)
	return d.RowsAffected
}
