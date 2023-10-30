package models

type Name struct {
	FirstName  string
	MiddleName string
	LastName   string
}

type Address struct {
	DoorNo     string
	StreetName string
	City       string
}

type Date struct {
	Day   int8
	Month int8
	Year  int32
}

type Customer struct {
	AccountNumber    int32
	FullName         Name
	Email            string
	ContactNo        int64
	PermanentAddress Address
	TemporaryAddress Address
	Password         string
	DOB              Date
}
