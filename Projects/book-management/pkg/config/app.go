package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDatabase(){
	database,err:= gorm.Open("mysql","Kritika:123456/book-management?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}

	database.AutoMigrate(&Book{})

	db = database

}