package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm" 
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	user := "root"
	pass := "Wanderlust@26"
	host := "localhost"
	port := "3306"
	dbname := "task_management"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open("mysql", URL) 
	if err != nil {
		panic(err.Error())
	}
	DB = db 

	return db,nil
}