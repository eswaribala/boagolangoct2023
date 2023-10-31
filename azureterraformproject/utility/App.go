package main

import (
	"azureterraformproject/models"
	"fmt"
	"reflect"
)

func main() {

	regions := [...]string{"US-East", "US-West", "Australia", "India"}

	//create the instance

	vmInstance1 := models.VMConfiguration{
		CPUCount:     3,
		RAM:          "32GB",
		Regions:      regions[0],
		AccessKey:    "Aegwfu3624586235",
		SecretKey:    "BIGIG*86869769",
		InstanceType: "t2.nano",
		Provider:     "azure",
		VMName:       "BOA_VM1",
	}

	fmt.Printf("VM instance Details%+v\n", vmInstance1)
	//check the data type

	typeOfArray := reflect.TypeOf(regions)
	fmt.Printf("Type of the Array=%s\n", typeOfArray)

	//check the size of the structure

	size := reflect.TypeOf(vmInstance1).Size()
	fmt.Printf("Size of the Struture=%d\n", size)

}
