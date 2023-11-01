package main

import (
	"azureterraformproject/interfaces"
	"azureterraformproject/models"
)

var (
	vmInterfaceRef interfaces.VMInterface
)

func main() {
	regions := [...]string{"US-East", "US-West", "Australia", "India"}
	//step1
	vmInstance := models.VMConfiguration{
		CPUCount:     3,
		RAM:          "32GB",
		Regions:      regions[0],
		AccessKey:    "Aegwfu3624586235",
		SecretKey:    "BIGIG*86869769",
		InstanceType: "t2.nano",
		Provider:     "azure",
		VMName:       "BOA_VM1",
	}

	//step2
	//Associate the instance to interface
	vmInterfaceRef = &vmInstance
	instanceName := "BOA_VM1"
	
	//step3
	//method invocation
	vmInterfaceRef.Start(&instanceName)
	vmInterfaceRef.Stop(true)
	vmInterfaceRef.Terminate(true)

}
