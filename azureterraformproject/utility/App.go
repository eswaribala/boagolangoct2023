package main

import "azureterraformproject/models"

func main() {

	regions := [...]string{"US-East", "US-West", "Australia", "India"}

	//create the instance

	vmInstance1 := models.VMConfiguration{
		CPUCount: 3,
		RAM:      "32GB",
		Regions:  regions,
	}

}
