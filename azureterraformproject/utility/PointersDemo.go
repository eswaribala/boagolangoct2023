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
}

func changeName(vmName *string) {
	*vmName = *vmName + strconv.Itoa(rand.Intn(10000000000000))
}
