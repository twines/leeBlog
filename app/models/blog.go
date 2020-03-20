package models

type Blog struct {
	Model
	Title       string
	Content     string
	Description string
	Good        int
	View        int
	UserId      int
	CategoryId  int
}
