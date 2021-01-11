package bootstrap

import (
	"goblog2/app/models/article"
	"goblog2/app/models/user"
	"goblog2/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	db := model.ContentDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &article.Article{})
}
