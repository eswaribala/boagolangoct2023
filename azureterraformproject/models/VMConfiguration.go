package models

type VMConfiguration struct {
	Provider     string `json:"provider"`
	VMName       string `json:"vm_name"`
	AccessKey    string `json:"access_key"`
	SecretKey    string `json:"secret_key"`
	Regions      string `json:"regions"`
	InstanceType string `json:"instance_type"`
	CPUCount     int    `json:"cpu_count"`
	RAM          string `json:"ram"`
}

type VMInstances struct {
	VMInstances []VMConfiguration `json:"instances"`
}
