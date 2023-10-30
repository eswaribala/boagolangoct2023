package main

import (
	"fmt"
	"strconv"
)

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
	var status bool
	for index, value := range loandecision {
		fmt.Println(index, value)
		for idx, val := range value {
			fmt.Println(idx, val)
		}
		//print loan approval status
		cibilScore, _ := strconv.Atoi(value[3])
		switch value[0] {
		case "young":
			if value[1] == "true" && value[2] == "true" && cibilScore > 700 {
				status = true
			}

		case "middle":
			if value[1] == "true" && value[2] == "true" && cibilScore > 700 {
				status = true
			}
		case "old":
			if value[1] == "true" && value[2] == "true" && cibilScore > 700 {
				status = true
			}
		}
	}
}
