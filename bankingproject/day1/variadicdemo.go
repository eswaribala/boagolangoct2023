package main

import (
	"errors"
	"fmt"
)

func main() {

	result, _ := fetchTransactions(2367264, 5000, 6000, 8000, 7900)
	fmt.Println(result)
	result, _ = fetchTransactions(2367265, 5000)
	fmt.Println(result)
	result, _ = fetchTransactions(2367265)
	fmt.Println(result)
}

func fetchTransactions(accountNumber int64, transactionAmount ...int32) (int32, error) {
	sum := 0
	if len(transactionAmount) == 0 {

		return 0, errors.New("no transaction happened")

	} else {
		for _, value := range transactionAmount {
			sum = int(int32(sum) + value)
		}
		return int32(sum), nil
	}

}
