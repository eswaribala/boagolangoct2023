package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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
	changeName(pvmName)
	fmt.Printf("After changing VM Name=%s\n", vmName)
	fmt.Printf("After Changing VM Name=%v\n", *pvmName)
	fmt.Printf("After Changing Address of VM Name=%v\n", pvmName)
	result := changeNameCallByValue(vmName)
	fmt.Printf("After changing VM Name using call by value=%s\n", result)
}

// call by reference
func changeName(vmName *string) {
	*vmName = *vmName + strconv.Itoa(rand.Intn(10000000000000))
}

// call by value
func changeNameCallByValue(vmName string) (result string) {
	vmName1 := vmName + strconv.Itoa(rand.Intn(100))
	fmt.Printf("Modified Value=%s\n", vmName1)
	return vmName1
}
