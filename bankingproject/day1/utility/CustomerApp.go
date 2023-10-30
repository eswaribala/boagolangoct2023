package main

import (
	"bankingproject/day1/models"
	"fmt"
	"math/rand"
)

func main() {

	//create the instance

	customerObj := models.Customer{
		AccountNumber: rand.Int31n(1000000000),
		FullName: models.Name{
			FirstName:  "Parameswari",
			MiddleName: "",
			LastName:   "Bala",
		},
		Email:     "Parameswaribala@gmail.com",
		ContactNo: 9952032862,
		PermanentAddress: models.Address{
			DoorNo:     "16",
			StreetName: "First Street",
			City:       "Chennai",
		},
		TemporaryAddress: models.Address{
			DoorNo:     "16",
			StreetName: "First Street",
			City:       "Chennai",
		},
		DOB: models.Date{
			Day:   2,
			Month: 12,
			Year:  1970,
		},
		Password: "Test@123",
	}

	fmt.Printf("Customer Details%+v\n", customerObj)
}
