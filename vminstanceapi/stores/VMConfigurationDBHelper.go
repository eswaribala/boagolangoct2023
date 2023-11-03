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
// SaveVMInstance godoc
// @Summary Save a new VM Instance
// @Description Save a new vminstance with the input payload
// @Tags vminstances
// @Accept  json
// @Produce  json
// @Param vmInstance body VMConfiguration true "Create vmInstance"
// @Success 200 {object} VMConfiguration
// @Router /vminstances/v1.0 [post]
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
// @Summary Get details of all vmInstances
// @Description Get details of all vmInstances
// @Tags vminstances
// @Accept  json
// @Produce  json
// @Success 200 {array} VMConfiguration
// @Router /vminstances/v1.0 [get]
func GetAllVMInstances(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	var vmInstances []VMConfiguration
	//db.Find(&customers)
	db.Find(&vmInstances)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstances)

}

// GetVMConfigurationByName select specific row
// @Summary Get details of  vmInstance by name
// @Description Get details of vmInstance by name
// @Tags vminstances
// @Param vmName path string true "VMName of the VMInstance"
// @Accept  json
// @Produce  json
// @Success 200 {object} VMConfiguration
// @Router /vminstances/v1.0/{vmName} [get]
func GetVMConfigurationByName(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	var vmInstance VMConfiguration
	params := mux.Vars(request)
	vmName := params["vmName"]
	log.Println("VM Name", vmName)
	db.Where("vm_name=?", vmName).First(&vmInstance)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(vmInstance)

}

// DeleteVMInstance godoc
// @Summary Delete vmInstance
// @Description Delete vmInstance by name
// @Tags vminstances
// @Accept  json
// @Produce  json
// @Param vmName path string true "VMName of the VMConfiguration"
// @Success 200 {object} VMConfiguration
// @Router /vminstances/v1.0/{vmName} [delete]
func DeleteVMInstance(response http.ResponseWriter, request *http.Request) {
	db, _ := CreateDBConnection()
	params := mux.Vars(request)
	vmName := params["vmName"]
	db.Where("vm_name=?", vmName).Delete(&VMConfiguration{})
	response.WriteHeader(http.StatusNoContent)
}

// UpdateVMInstance godoc
// @Summary Update existing vmInstance
// @Description Update existing vmInstance with the input vmInstance
// @Tags vminstances
// @Accept  json
// @Produce  json
// @Param vmInstance body VMConfiguration true "Create customer"
// @Success 200 {object} VMConfiguration
// @Router /vminstances/v1.0 [put]
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
