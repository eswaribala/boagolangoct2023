package main

import (
	"azureterraformproject/models"
	"encoding/json"
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

	result, _ := json.Marshal(vmInstance1)
	fmt.Printf("The Data=%s", result)

}
