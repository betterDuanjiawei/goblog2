package user

import (
	"goblog2/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(150);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(150);unique; vaild:"email"`
	Password        string `gorm:"type:varchar(150)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
