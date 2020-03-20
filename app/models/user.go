package models

type User struct {
	Model
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	QQ       string `json:"qq"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Gender   uint8  `json:"gender"`
	Status   uint8  `json:"status"`
	Address  string `json:"address"`
}
