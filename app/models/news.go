package models

type News struct {
	Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	Keyword     string `json:"keyword"`
	Status      uint8  `json:"status"`
	Description string `json:"description"`
}
