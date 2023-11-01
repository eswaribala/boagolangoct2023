package interfaces

type VMInterface interface {
	ConnectionHelper()
	Start(vmInstance *string)
	Stop(start bool)
	Terminate(stop bool)
}
