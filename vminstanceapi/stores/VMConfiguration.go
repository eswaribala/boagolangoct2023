package stores

type Subnet struct {
	CIDR             string
	VPCID            string
	ZONEID           string
	AvailabilityZone bool
}

type VMConfiguration struct {
	Provider     string `json:"provider"`
	VMName       string `json:"vm_name" gorm:"primary_key"`
	AccessKey    string `json:"access_key"`
	SecretKey    string `json:"secret_key"`
	Regions      string `json:"regions"`
	InstanceType string `json:"instance_type"`
	CPUCount     int    `json:"cpu_count"`
	RAM          string `json:"ram"`
	//SubnetType   Subnet //has relationship
}
