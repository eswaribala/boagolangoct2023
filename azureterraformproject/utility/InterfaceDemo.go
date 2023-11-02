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
		CPUCount:     4,
		RAM:          "64GB",
		Regions:      regions[0],
		AccessKey:    "Aegwfu3624586235",
		SecretKey:    "BIGIG*86869769",
		InstanceType: "t2.nano",
		Provider:     "azure",
		VMName:       "BOA_VM2",
	}

	//step2
	//Associate the instance to interface
	vmInterfaceRef = &vmInstance
	instanceName := "BOA_VM2"

	//step3
	//method invocation
	vmInterfaceRef.ConnectionHelper()
	vmInterfaceRef.SaveInstance(&vmInstance)
	vmInterfaceRef.RetrieveAllInstances()
	vmInterfaceRef.RetrieveInstanceByName(&instanceName)
	vmInterfaceRef.Start(&instanceName)
	vmInterfaceRef.Stop(true)
	vmInterfaceRef.Terminate(true)

}
