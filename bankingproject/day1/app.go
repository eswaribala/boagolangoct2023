package main

import (
	"fmt"
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
	fmt.Printf("Customre Name=%s having accountNo=%d", name, accountNumber)
}
