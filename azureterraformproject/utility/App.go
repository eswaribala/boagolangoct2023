package main

import (
	"azureterraformproject/models"
	"fmt"
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

	fmt.Printf("VM instance Details%+v", vmInstance1)

}
