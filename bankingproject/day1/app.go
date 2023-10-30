package main

import (
	"fmt"
	"os"
	"strconv"
)

func init() {
	fmt.Println("Initialize Mysql Server!!!")
}

var (
	accountNumber int64
	name          string
	contactNo     int64
	email         string
	password      string
	address       string
)

func main() {
	accountNumber = 487687468
	name = "Parameswari"
	password = "Test@123"
	contactNo = 9952032862
	email = "parameswaribala@gmail.com"
	address = "chennai"
	fmt.Println("Rocking with Go!!!")
	fmt.Printf("Customre Name=%s having accountNo=%d\n", name, accountNumber)
	//invoking function from external go file in the same package
	//helper()
	//read command line arguments
	args := os.Args
	//find out the number of args passed
	fmt.Printf("The length of the arguments = %d\n", len(args))
	if len(args) > 2 {
		data, err := readServerProperties(args)
		if err != nil {
			fmt.Println(err)
		} else {
			for index, value := range data {
				fmt.Println(index, value)
			}
		}
	}
}

func readServerProperties(properties []string) ([]string, error) {

	//create the array
	serverProperties := make([]string, 4)

	//traditional loop
	/*
		for i := 1; i < len(properties); i++ {
			fmt.Println(properties[i])
		}
	*/

	port, err := strconv.Atoi(properties[2])

	if err != nil {
		return nil, err
	} else {

		if port > 1027 {

			//enhanced for loop
			for index, value := range properties {
				serverProperties[index] = value
			}
		} else {
			return nil, err
		}
	}
	return serverProperties, nil
}
