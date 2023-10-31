package main

import (
	"azureterraformproject/models"
	"fmt"
	"math/rand"
	"strconv"
)

//generate 100 vm instances

func main() {
	//declared
	vmInstances := make([]models.VMConfiguration, 100)
	regions := [...]string{"US-East", "US-West", "Australia", "India"}
	insatnceTypes := [...]string{"nano", "micro", "medium", "large", "xlarge"}
	for i := 0; i < len(vmInstances); i++ {
		vmInstances[i] = models.VMConfiguration{
			CPUCount:     rand.Intn(16),
			RAM:          strconv.Itoa(rand.Intn(100)) + "GB",
			Regions:      regions[0],
			AccessKey:    "Aegwfu3624586235",
			SecretKey:    "BIGIG*86869769",
			InstanceType: insatnceTypes[rand.Intn(4)],
			Provider:     "azure",
			VMName:       "BOA_VM" + strconv.Itoa(i),
		}
	}

	for _, value := range vmInstances {
		fmt.Printf("Vm Instance Details %+v\n", value)
	}

}
