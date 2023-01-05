package database

import (
	"fmt"

	"absensi-api.com/config"
	"absensi-api.com/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&model.Activity{}); err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&model.Attendance{}); err != nil {
		panic(err)
	}

	return db
}
