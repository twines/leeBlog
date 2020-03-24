package models

type Site struct {
	Model
	SiteName    string `json:"siteName"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Number      string `json:"number"`
	Status      uint8  `json:"status"`
	Logo        string `json:"logo"`
}
