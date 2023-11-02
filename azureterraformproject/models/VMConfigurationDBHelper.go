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

// SaveVMInstance insert the data into table
func SaveVMInstance(vmInstance *VMConfiguration) {

	db := CreateDBConnection()

	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	db.Create(vmInstance)
	tx.Commit()

}

// GetAllVMInstances select all rows
func GetAllVMInstances() (vmInstancesResult *[]VMConfiguration) {
	db := CreateDBConnection()

	var vmInstances []VMConfiguration
	//db.Find(&customers)
	db.Find(&vmInstances)
	return &vmInstances
}

func GetVMConfigurationByName(vmName *string) (vmInstancesResult *VMConfiguration) {
	db := CreateDBConnection()

	var vmInstance VMConfiguration

	db.First(&vmInstance, vmName)

	return &vmInstance

}
