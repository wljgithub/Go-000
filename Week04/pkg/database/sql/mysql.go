package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql()*gorm.DB  {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
	"root",
		"root",
		"localhost:3306",
		"mydb",
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	return db
}