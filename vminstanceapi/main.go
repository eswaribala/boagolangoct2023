package main

import (
	"github.com/gorilla/mux"
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

}
