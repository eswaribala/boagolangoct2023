package facades

import "vminstanceapi/stores"

type VMInterface interface {
	ConnectionHelper()
	SaveInstance(vmInstance *stores.VMConfiguration)
	RetrieveAllInstances() (vmInstances *[]stores.VMConfiguration)
	RetrieveInstanceByName(name *string) (vmInstances *stores.VMConfiguration)
	Start(vmInstance *string)
	Stop(start bool)
	Terminate(stop bool)
}
