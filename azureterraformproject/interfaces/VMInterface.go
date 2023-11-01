package interfaces

type VMInterface interface {
	Start(vmInstance *string)
	Stop(start bool)
	Terminate(stop bool)
}
