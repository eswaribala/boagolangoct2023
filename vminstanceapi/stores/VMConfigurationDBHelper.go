package stores

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
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
func SaveVMInstance(response http.ResponseWriter, request *http.Request) {

	db, _ := CreateDBConnection()
	var vmInstance VMConfiguration
	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	json.NewDecoder(request.Body).Decode(&vmInstance)
	db.Create(vmInstance)
	tx.Commit()
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstance)
}

// GetAllVMInstances select all rows
func GetAllVMInstances(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	var vmInstances []VMConfiguration
	//db.Find(&customers)
	db.Find(&vmInstances)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstances)

}

func GetVMConfigurationByName(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	var vmInstance VMConfiguration
	params := mux.Vars(request)
	vmName := params["vmName"]
	db.First(&vmInstance, vmName)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstance)

}

func DeleteVMInstance(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	params := mux.Vars(request)
	vmName := params["vmName"]
	db.Where("VMName=?", vmName).Delete(&VMConfiguration{})
	response.WriteHeader(http.StatusOK)
}

func UpdateVMInstance(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	var vmInstance VMConfiguration
	tx := db.Begin()
	db.Set("gorm:auto_preload", true)
	json.NewDecoder(request.Body).Decode(&vmInstance)
	db.Save(vmInstance)
	tx.Commit()
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstance)
}
