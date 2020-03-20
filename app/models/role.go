package models

type Role struct {
	Model
	RoleName    string       `json:"roleName"`
	Description string       `json:"description"`
	Permission  []Permission `json:"permission" gorm:"many2many:role_permissions;association_autoupdate:false;association_autocreate:false"`
}
