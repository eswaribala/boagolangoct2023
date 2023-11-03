package stores

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	Config "vminstanceapi/config"
)

func handlePanicScenario() {

	value := recover()
	fmt.Println("Recovered from Panic", value)
}

func CreateDBConnection() (*gorm.DB, error) {
	//db, err := gorm.Open("mysql", "root:vignesh@(localhost:3306)/azuredb?parseTime=true")

	defer handlePanicScenario()
	if len(Config.DbURL(Config.BuildDBConfig())) == 0 {
		panic("Connection string is empty")
		//return nil, errors.New("Connection String Empty")
	} else {

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
		return db, nil
	}

}

// SaveVMInstance insert the data into table
func SaveVMInstance(vmInstance *VMConfiguration) {

	db, _ := CreateDBConnection()

	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	db.Create(vmInstance)
	tx.Commit()

}

// GetAllVMInstances select all rows
func GetAllVMInstances() (vmInstancesResult *[]VMConfiguration) {
	db, _ := CreateDBConnection()

	var vmInstances []VMConfiguration
	//db.Find(&customers)
	db.Find(&vmInstances)
	return &vmInstances
}

func GetVMConfigurationByName(vmName *string) (vmInstancesResult *VMConfiguration) {
	db, _ := CreateDBConnection()

	var vmInstance VMConfiguration

	db.First(&vmInstance, vmName)

	return &vmInstance

}
