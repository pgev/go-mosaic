package service

import (
	"errors"
	"sync"
)

var (
	// ErrAlreadyRunning is returned on a call to run an already running service.
	ErrAlreadyRunning = errors.New("already started")
)

// Servicable provides Start() and Stop() interfaces to be implemented to become
// eligible to act as a service.
type Servicable interface {
	Start() error
	Stop() error
}

// Service defines a service that can be started and stopped.
type Service interface {
	// Start the service.
	Start() error

	// Stop the service.
	Stop() error

	// IsRunning returns true when the service is running.
	IsRunning() bool

	// Wait blocks until the service is stopped.
	Wait()

	// String representation of the service.
	String() string
}

type service struct {
	name      string
	isRunning bool
	quit      chan struct{}
	mux       sync.Mutex

	impl Servicable
}

// IsRunning returns true when the service is running.
func (bs *BaseService) IsRunning() bool {
	return bs.isRunning
}

// Start the service.
func (bs *BaseService) Start() error {
	bs.mux.Lock()
	defer bs.mux.Unlock()
	if bs.isRunning {
		return ErrAlreadyRunning
	}

	if err := bs.impl.Start(); err != nil {
		return err
	}
	bs.isRunning = true
	bs.quit = make(chan struct{})

	return nil
}

// Stop the service.
func (bs *BaseService) Stop() error {
	bs.mux.Lock()
	defer bs.mux.Unlock()
	if !bs.isRunning {
		return nil
	}

	if err := bs.impl.Stop(); err != nil {
		return err
	}
	bs.isRunning = false
	close(bs.quit)

	return nil
}

// Wait blocks until the service is stopped.
func (bs *BaseService) Wait() {
	<-bs.quit
}

// String returns a string representation of the service.
func (bs *BaseService) String() string {
	return bs.name
}
