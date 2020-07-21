package service


// Servicable provides OnStart() and OnStop() interfaces to be implemented
// to become eligible to act as a service.
type Servicable interface {
	// OnStart is called by ServiceImpl when the Service is started.
	// An error is returned to be thrown by Start()
	OnStart() error

	// OnStop is called by ServiceImpl when the Service is stopped.
	// No error should be thrown (TODO: re-evaluate later)
	OnStop()
}

// Service defines a service that can be started, waited and stopped.
type Service interface {
	// Start the service.
	Start() error

	// Stop the service.
	Stop()

	// IsRunning returns true when the service is running.
	IsRunning() bool

	// Wait blocks until the service is stopped.
	Wait()

	// String representation of the service.
	String() string
}
