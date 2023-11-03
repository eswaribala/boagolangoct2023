package main

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "vminstanceapi/docs"
	"vminstanceapi/stores"
)

// @title VMConfiguration API
// @version 1.0
// @description This is api service for managing VM Configuration
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7072
// @BasePath /

func main() {

	router := mux.NewRouter()

	//create api endpoints
	router.HandleFunc("/vminstances/v1.0", stores.SaveVMInstance).Methods("POST")
	router.HandleFunc("/vminstances/v1.0", stores.GetAllVMInstances).Methods("GET")
	router.HandleFunc("/vminstances/v1.0/{vmName}", stores.GetVMConfigurationByName).Methods("GET")
	router.HandleFunc("/vminstances/v1.0/{vmName}", stores.DeleteVMInstance).Methods("DELETE")
	router.HandleFunc("/vminstances/v1.0", stores.UpdateVMInstance).Methods("PUT")

	stores.CreateDBConnection()
	//swagger documentation
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//api deployment
	log.Fatal(http.ListenAndServe("localhost:7072", router))

}
