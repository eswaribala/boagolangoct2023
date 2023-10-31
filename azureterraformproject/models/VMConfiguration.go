package models

type VMConfiguration struct {
	Provider     string
	VMName       string
	AccessKey    string
	SecretKey    string
	Regions      string
	InstanceType string
	CPUCount     int
	RAM          string
}
