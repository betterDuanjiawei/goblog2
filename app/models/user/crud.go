package user

import (
	"goblog2/pkg/logger"
	"goblog2/pkg/model"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
