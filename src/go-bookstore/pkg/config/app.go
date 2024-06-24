package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	// Update with your credentials
	dsn := "Mayank123:Mayank@199929@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		// Log the error message with more detail
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
