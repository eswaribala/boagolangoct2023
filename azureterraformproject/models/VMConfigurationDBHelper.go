package models

import (
	Config "azureterraformproject/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"log"
)

func CreateDBConnection() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:vignesh@(localhost:3306)/?parseTime=true")
	db, err := gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	//db.Exec("Create Database azuredb")
	db.Exec("use azuredb")
	//generate table
	//whatever struct will be converted to table
	db.AutoMigrate(&VMConfiguration{})
	return db

}
