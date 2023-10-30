package main

import (
	"fmt"
	"math/rand"
)

func main() {
	accountNumbers := make([]int32, 100)

	for i := 0; i < len(accountNumbers); i++ {
		accountNumbers[i] = rand.Int31n(1000000000)
	}

	//print account numbers
	for index, value := range accountNumbers {
		fmt.Println(index, value)
	}

}
