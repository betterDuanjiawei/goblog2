package bootstrap

import (
	"goblog2/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ContentDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
