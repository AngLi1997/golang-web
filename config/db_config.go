package config

import (
	"go-test/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//func ConnectToMysqlDB(username string, password string, host string, port int, databaseName string) {
//	dsn := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
//	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	global.DB = db
//}

func ConnectToSQLiteDB(dbPath string) {
	dsn := dbPath
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	global.DB = db
}
