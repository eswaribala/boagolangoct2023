package main

import (
	"fmt"
	"os"
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
		readServerProperties(args)
	}
}

func readServerProperties(properties []string) {

	//create the array

	//traditional loop
	/*
		for i := 1; i < len(properties); i++ {
			fmt.Println(properties[i])
		}
	*/

	//enhanced for loop
	for index, value := range properties {
		fmt.Println(index, value)
	}

}
