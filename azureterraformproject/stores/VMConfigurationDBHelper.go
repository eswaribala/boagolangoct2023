package stores

import (
	"azureterraformproject/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)
	- "github.com/jinzhu/gorm/dialects/mysql"
)

func createDBConnection() *gorm.DB{

	db, err := gorm.Open("mysql", "root:vignesh@(mysql:3306)/?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Exec("Create Database azuredb")
	db.Exec("use azuredb")
	//generate table
	//whatever struct will be converted to table
	db.AutoMigrate(&models.VMConfiguration{})

	return db


}
