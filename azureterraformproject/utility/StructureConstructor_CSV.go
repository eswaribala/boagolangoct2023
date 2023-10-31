package main

import (
	"azureterraformproject/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fileContent, err := os.Open("resources/data.csv")

	if err != nil {

		fmt.Println("Error occurred", err)
	} else {
		fmt.Println("File Successfully Opened")
		fileReader := csv.NewReader(fileContent)
		records, _ := fileReader.ReadAll()
		vmInstances := make([]models.VMConfiguration, len(records))

		for index, value := range records {
			if index == 0 {
				continue
			}

			vmInstances[index] = VMConfigurationV1(value)
			fmt.Printf("Vm Instance=%+v\n", vmInstances[index])
		}

	}

}

// VMConfigurationV1 simulating parameterized constructor
func VMConfigurationV1(data []string) models.VMConfiguration {

	cpuCount, _ := strconv.Atoi(data[6])
	vmInstance := models.VMConfiguration{
		CPUCount:     cpuCount,
		RAM:          data[7],
		Regions:      data[4],
		AccessKey:    data[2],
		SecretKey:    data[3],
		InstanceType: data[5],
		Provider:     data[0],
		VMName:       data[1],
	}
	return vmInstance

}
