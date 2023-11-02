package interfaces

import "azureterraformproject/models"

type VMInterface interface {
	ConnectionHelper()
	SaveInstance(vmInstance *models.VMConfiguration)
	RetrieveAllInstances() (vmInstances *[]models.VMConfiguration)
	RetrieveInstanceByName(name *string) (vmInstances *models.VMConfiguration)
	Start(vmInstance *string)
	Stop(start bool)
	Terminate(stop bool)
}
