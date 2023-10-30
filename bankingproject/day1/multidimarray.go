package main

import "fmt"

func main() {

	//declare 2d array
	loandecision := [][]string{}
	row1 := []string{"young", "true", "true", "750"}
	row2 := []string{"young", "true", "false", "800"}
	row3 := []string{"young", "false", "false", "0"}
	row4 := []string{"middle", "true", "false", "850"}
	row5 := []string{"old", "true", "true", "750"}

	loandecision = append(loandecision, row1)
	loandecision = append(loandecision, row2)
	loandecision = append(loandecision, row3)
	loandecision = append(loandecision, row4)
	loandecision = append(loandecision, row5)

	//show the table

	for i := range loandecision {
		fmt.Println(i)
	}
}
