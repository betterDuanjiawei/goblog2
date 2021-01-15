package user

import (
	"goblog2/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(175);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(175);unique;" valid:"email"`
	Password        string `gorm:"type:varchar(175)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
