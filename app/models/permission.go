package models

type Permission struct {
	Model
	Path     string `json:"path"`
	Title    string `json:"title"`
	ApiPath  string `json:"apiPath"`
	ParentId uint   `json:"parentId" gorm:"default:0"`
	Icon     string `json:"icon"`
}
