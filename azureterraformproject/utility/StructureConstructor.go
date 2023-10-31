package main

import (
	"azureterraformproject/models"
	"fmt"
)

var (
	provider     string
	vmname       string
	accessKey    string
	secretKey    string
	regions      string
	instanceType string
	cpuCount     int
	ram          string
)

func main() {
	//dynamic values
	fmt.Println("Enter Provider Name")
	fmt.Scanln(&provider)
	fmt.Println("Enter VM Name")
	fmt.Scanln(&vmname)
	fmt.Println("Enter Access Key")
	fmt.Scanln(&accessKey)
	fmt.Println("Enter Secret Key")
	fmt.Scanln(&secretKey)
	fmt.Println("Enter Region")
	fmt.Scanln(&regions)
	fmt.Println("Enter Instance Type")
	fmt.Scanln(&instanceType)
	fmt.Println("Enter CPU Count")
	fmt.Scanln(&cpuCount)
	fmt.Println("Enter RAM")
	fmt.Scanln(&ram)

	vmInstance := VMConfiguration(provider, vmname, accessKey, secretKey, regions, instanceType, cpuCount, ram)
	fmt.Printf("VM Instance%+v", vmInstance)

}

// VMConfiguration simulating parameterized constructor
func VMConfiguration(provider string, vmName string, accessKey string, secretKey string,
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
