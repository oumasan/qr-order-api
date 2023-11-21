package models

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := "user:password@tcp(mysql:3306)/qr_order_system?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
}

func getDB() *gorm.DB {
	return DB
}