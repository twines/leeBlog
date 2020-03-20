package models

type Category struct {
	Model
	CategoryName string
	ParentId     int
	Keyword      string
	Description  string
	Title        string
}
