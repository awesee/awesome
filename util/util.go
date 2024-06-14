package util

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GormOpen(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	CheckErr(err)
	return db
}
