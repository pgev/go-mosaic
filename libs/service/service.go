package service

import (

)

// Service defines a service that can be started and stopped.
type Service interface {
	// Start the service
	Start() error
	StartImpl() error

	// Stop the service
	Stop() error
	StopImpl()

	// Reset() error

	// IsRunning returns true when the service is running
	IsRunning() bool

	// Quit returns a channel, which is closed once service is stopped.
	Quit() <-chan struct{}

	// String representation of the service
	String() string
}

// BaseService provides Start() and Stop()
type BaseService struct {
	name    string
	// set atomically
	started uint32
	stopped uint32
	quit    chan struct{}

	// The "subclass" of BaseService
	impl Service
}
