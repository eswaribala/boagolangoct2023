package main

import (
	"fmt"
	"strconv"
)

func main() {

	customers := make([]string, 4, 6)
	//confirm append beyond capacity???

	//The append built-in function appends elements to the end of a slice. If
	// it has sufficient capacity, the destination is resliced to accommodate the
	// new elements. If it does not, a new underlying array will be allocated.
	// Append returns the updated slice. It is therefore necessary to stores the
	// result of append, often in the variable holding the slice itself:
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	// As a special case, it is legal to append a string to a byte slice, like this:
	//	slice = append([]byte("hello "), "world"...)

	for i := 0; i < 10; i++ {
		customers = append(customers, "customer"+strconv.Itoa(i))
	}

	for index, value := range customers {
		fmt.Println(index, value)
	}

}
