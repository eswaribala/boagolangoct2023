package main

import (
	"azureterraformproject/models"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	fileContent, err := os.Open("resources/data.csv")

	if err != nil {

		fmt.Println("Error occurred", err)
	} else {
		fmt.Println("File Successfully Opened")
	}

	fileReader := csv.NewReader(fileContent)
	records, _ := fileReader.ReadAll()
	fmt.Println(records)

	//vmInstance := VMConfiguration(provider, vmname, accessKey, secretKey, regions, instanceType, cpuCount, ram)
	//fmt.Printf("VM Instance%+v", vmInstance)

}

// VMConfigurationV1 simulating parameterized constructor
func VMConfigurationV1(provider string, vmName string, accessKey string, secretKey string,
	regions string, instanceType string, cpuCount int, ram string) *models.VMConfiguration {

	vmInstance := models.VMConfiguration{
		CPUCount:     cpuCount,
		RAM:          ram,
		Regions:      regions,
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		InstanceType: instanceType,
		Provider:     provider,
		VMName:       vmName,
	}
	return &vmInstance

}
