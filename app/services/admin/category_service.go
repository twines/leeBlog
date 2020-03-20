package admin

import "leeblog.com/app/models"

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}
func (c *CategoryService) GetCategoryList(page int, limit int) (categorySlice []models.Category, count int) {
	models.DB().Offset((page - 1) * page).Limit(limit).Find(&categorySlice)
	models.DB().Model(&models.Category{}).Count(&count)
	return
}
