package user

import (
	"goblog2/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(150);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(150);default:NULL;unique;"`
	Password string `gorm:"column:password;type:varchar(150)"`
}
