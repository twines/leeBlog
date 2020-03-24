package admin

import "leeblog.com/app/models"

type SiteService struct {
}

func NewSiteService() *SiteService {
	return &SiteService{}
}
func (s *SiteService) AddSiteConfig(site models.Site) {
	models.DB().Unscoped().Delete(&models.Site{})
	models.DB().FirstOrCreate(&site)
}
func (s *SiteService) GetSiteConfig() (site models.Site) {
	models.DB().First(&site)
	return
}
