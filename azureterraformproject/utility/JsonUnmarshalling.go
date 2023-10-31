package main

import (
	"azureterraformproject/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	vmInstances models.VMInstances
)

func main() {

	jsonFile, err := os.Open("resources/tags.json")
	if err != nil {

		fmt.Println("Error occurred", err)
	} else {
		fmt.Println("File Successfully Opened")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &vmInstances)
		for key, value := range vmInstances.VMInstances {
			fmt.Println(key, value)

		}

	}

}
