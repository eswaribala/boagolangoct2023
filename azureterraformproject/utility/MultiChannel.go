package main

import (
	"azureterraformproject/models"
	"github.com/gruntwork-io/terratest/modules/azure"
	"log"
)

func main() {

	awsChannel := make(chan string)
	gcpChannel := make(chan int)
	azureChannel := make(chan models.VMConfiguration)

	go func(pChannel chan int) {
		pChannel <- 1

	}(gcpChannel)

	go AWSCloud(awsChannel)
	go AzureCloud(azureChannel)

}

func AWSCloud(pChannel chan string) {

	pChannel <- "Aws cloud sending message"
}

func AzureCloud(pChannel chan models.VMConfiguration) {

	_, err := azure.AppExistsE("testapp", "vebgroup2022", "89e10581-ca05-46fa-af94-9ff90fce5750")

	if err != nil {
		log.Fatal("Error in finding resource", err)
	} else {

		vmInstance1 := models.VMConfiguration{
			CPUCount:     3,
			RAM:          "32GB",
			Regions:      "India",
			AccessKey:    "Aegwfu3624586235",
			SecretKey:    "BIGIG*86869769",
			InstanceType: "t2.nano",
			Provider:     "azure",
			VMName:       "BOA_VM1",
		}

		pChannel <- vmInstance1
	}
}
