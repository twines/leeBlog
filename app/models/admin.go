package models

type Admin struct {
	Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Status   uint8  `json:"status" gorm:"default:1"`
	QQ       string `json:"qq"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	TrueName string `json:"trueName"`
	RoleId   uint   `json:"roleId"`
	Role     Role   `json:"role"`
}
