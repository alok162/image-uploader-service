package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() {
	DB, err := gorm.Open("mysql", "root@(localhost)/image_service_db?parseTime=true")
	if err != nil {
		fmt.Println("failed to connect database", err)
		return
	}
	fmt.Println("successfullt db got connected", DB)
	setUpDbConnection(DB)
}

func setUpDbConnection(db *gorm.DB) {
	DB = db
}

func GetDBInstance() *gorm.DB {
	return DB
}
