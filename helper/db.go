package helper

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:password@tcp(localhost:3306)/devops?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Seems to have worked to create the db connection? How...")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
