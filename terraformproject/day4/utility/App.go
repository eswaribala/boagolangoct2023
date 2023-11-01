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

	vmInstance2 := models.VMConfiguration{
		CPUCount:     13,
		RAM:          "64GB",
		Regions:      regions[2],
		AccessKey:    "Aegwfu3624586235546356",
		SecretKey:    "BIGIG*868697694365467",
		InstanceType: "t2.micro",
		Provider:     "azure",
		VMName:       "BOA_VM2",
	}

	result := reflect.DeepEqual(vmInstance2, vmInstance1)
	fmt.Println("Is destination equal to Source", result)

	//type comparison
	structType := reflect.TypeOf(vmInstance1)
	fmt.Println(structType)

	if structType.Kind() == reflect.TypeOf(vmInstance1).Kind() {
		fmt.Println("They are same")
	}

	//any data type
	vnetData := []any{"VNET1", 8080, true}

	//show only numerical data
	for _, value := range vnetData {
		if reflect.TypeOf(value).Kind() == reflect.Int {
			fmt.Println(value)
		}
	}
}
