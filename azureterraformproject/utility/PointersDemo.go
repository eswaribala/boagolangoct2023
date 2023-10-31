package main

import "fmt"

var (
	vmName string
	//declare pointer
	pvmName *string
)

// change vm name using function
func main() {
	vmName = "BOA-VM1"
	//assign address of vmName to pointer
	pvmName = &vmName

	//display the vmName
	fmt.Printf("VM Name=%s\n", vmName)
	fmt.Printf("Before Changing VM Name=%v\n", *pvmName)
	fmt.Printf("Before Changing Address of VM Name=%v\n", pvmName)

}
