package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"terraformproject/day4/models"
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
		//mapping json data golang object
		json.Unmarshal(byteValue, &vmInstances)
		for key, value := range vmInstances.VMInstances {
			fmt.Println(key, value)

		}

	}

}