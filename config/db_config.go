package config

import (
	"go-test/global"
	"go-test/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToSQLiteDB(dbPath string, autoMigrate bool) {
	dsn := dbPath
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	global.DB = db
	if autoMigrate {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			return
		}
	}
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
