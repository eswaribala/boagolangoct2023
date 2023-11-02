package models

import (
	"log"
)

type Subnet struct {
	CIDR             string
	VPCID            string
	ZONEID           string
	AvailabilityZone bool
}

type VMConfiguration struct {
	Provider     string `json:"provider"`
	VMName       string `json:"vm_name"`
	AccessKey    string `json:"access_key"`
	SecretKey    string `json:"secret_key"`
	Regions      string `json:"regions"`
	InstanceType string `json:"instance_type"`
	CPUCount     int    `json:"cpu_count"`
	RAM          string `json:"ram"`
	//SubnetType   Subnet //has relationship
}

type VMInstances struct {
	VMInstances []VMConfiguration `json:"instances"`
}

// interface implementation
// interface method

func (vmConfiguration *VMConfiguration) ConnectionHelper() {
	CreateDBConnection()
}

func (vmConfiguration *VMConfiguration) SaveInstance(vmInstance *VMConfiguration) {

	SaveVMInstance(vmInstance)
}

func (vmConfiguration *VMConfiguration) RetrieveAllInstances() (vmInstances *[]VMConfiguration) {
	return GetAllVMInstances()
}
func (vmConfiguration *VMConfiguration) RetrieveInstanceByName(name *string) (vmInstances *VMConfiguration) {
	return GetVMConfigurationByName(name)
}

func (vmConfiguration *VMConfiguration) Start(vmInstanceName *string) {

	if vmConfiguration.VMName == *vmInstanceName {
		log.Println("Instance Started")
	}

}
func (vmConfiguration *VMConfiguration) Stop(start bool) {
	if start {
		log.Println("Instance Stopped")
	}

}
func (vmConfiguration *VMConfiguration) Terminate(stop bool) {

	if stop {
		log.Println("Instance Terminated")
	}
}
