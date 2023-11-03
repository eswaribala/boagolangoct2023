package main

import (
	"github.com/gorilla/mux"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"vminstanceapi/stores"
)

func main() {

	router := mux.NewRouter()

	//create api endpoints
	router.HandleFunc("/vminstances/v1.0", stores.SaveVMInstance).Methods("POST")
	router.HandleFunc("/vminstances/v1.0", stores.GetAllVMInstances).Methods("GET")
	router.HandleFunc("/vminstances/v1.0/{vmName}", stores.GetVMConfigurationByName).Methods("GET")
	router.HandleFunc("/vminstances/v1.0/{vmName}", stores.DeleteVMInstance).Methods("DELETE")
	router.HandleFunc("/vminstances/v1.0", stores.UpdateVMInstance).Methods("PUT")

	//swagger documentation

	//api deployment
	log.Fatal(http.ListenAndServe("7074", router))

}
