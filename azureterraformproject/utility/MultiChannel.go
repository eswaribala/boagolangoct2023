package main

import (
	"azureterraformproject/models"
	"fmt"
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

	func() {
		select {
		case m1 := <-awsChannel:
			fmt.Printf("AWS Message%s\n", m1)
		case m2 := <-azureChannel:
			fmt.Printf("Azure Message%+v\n", m2)
		case m3 := <-gcpChannel:
			fmt.Printf("GCP Message%d\n", m3)
		}
	}()

	fmt.Println("All Tasks Completed")
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
