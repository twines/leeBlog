package admin

import "leeblog.com/app/models"

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}
func (ns *NewsService) GetNewsList(page int, limit int) (newsSlice []models.News, count int) {
	models.DB().Order("id desc").Offset((page - 1) * limit).Limit(limit).Find(&newsSlice).Offset(-1).Limit(-1).Count(&count)
	return
}
func (ns *NewsService) GetNewsDetail(id int) (news models.News) {
	models.DB().First(&news, 1)
	return
}

func (ns *NewsService) DeleteNews(id int) bool {
	res := models.DB().Where("id=?", id).Delete(&models.News{})
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
func (ns *NewsService) UpdateNews(news models.News) bool {
	res := models.DB().Save(&news)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
func (ns *NewsService) AddNews(news models.News) bool {
	res := models.DB().Create(&news)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
