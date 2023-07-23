package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	databse, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/anama"))
	if err != nil {
		panic(err)
	}

	databse.AutoMigrate(&Product{})

	DB = databse

}
