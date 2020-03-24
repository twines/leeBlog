package models

// 用户留言
type Message struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content" gorm:"type:text"`
	Status  uint8  `json:"status" gorm:"default:0"`
}
