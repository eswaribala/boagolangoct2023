package main

import (
	"fmt"
	"strconv"
)

var (
	count int
)

func main() {

	tags := make(map[string]string)
	fmt.Println("Enter no of tags")
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		tags["tag"+strconv.Itoa(i)] = "Resource" + strconv.Itoa(i)

	}

	for key, value := range tags {
		fmt.Println(key, value)
	}

}