package web

import "leeblog.com/app/models"

type Category struct {
}

func NewCategory() *Category {
	return &Category{}
}

func (category *Category) GetChildCategory() (categorySlice []models.Category) {
	models.DB().Where("parent_id>?", 0).Find(&categorySlice)
	return
}
func (category *Category) GetParentCategory() (categorySlice []models.Category) {
	models.DB().Where("parent_id=?", 0).Find(&categorySlice)
	return
}
