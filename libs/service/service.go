package service

import (
	"errors"
	"sync"
)

var (
	// ErrAlreadyRunning is returned on a call to run an already running service.
	ErrAlreadyRunning = errors.New("already running")
)

// Servicable provides Start() and Stop() interfaces to be implemented to become
// eligible to act as a service.
type Servicable interface {
	// Start does not need to be a thread safe. Service implementation will
	// take care of it.
	Start() error

	// Stop does not need to be a thread safe. Service implementation will
	// take care of it.
	Stop() error
}

// Service defines a service that can be started, waited and stopped.
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

// NewService creates a new service from a servicable object.
func NewService(name string, impl Servicable) *Service {
	return &service{
		name:      name,
		isRunning: false,
		impl:      impl,
	}
}

// IsRunning returns true when the service is running.
func (bs *service) IsRunning() bool {
	return bs.isRunning
}

// Start the service.
func (bs *service) Start() error {
	bs.mux.Lock()
	defer bs.mux.Unlock()
	if bs.isRunning {
		return ErrAlreadyRunning
	}

	if err := bs.impl.Start(); err != nil {
		return err
	}
	bs.quit = make(chan struct{})
	bs.isRunning = true

	return nil
}

// Stop the service.
func (bs *service) Stop() error {
	bs.mux.Lock()
	defer bs.mux.Unlock()
	if !bs.isRunning {
		return nil
	}

	if err := bs.impl.Stop(); err != nil {
		return err
	}
	close(bs.quit)
	bs.isRunning = false

	return nil
}

// Wait blocks until the service is stopped.
func (bs *serbice) Wait() {
	bs.mux.Lock()
	defer bs.mux.Unlock()
	if !bs.isRunning {
		return
	}
	bs.mux.Unlock()

	<-bs.quit
}

// String returns a string representation of the service.
func (bs *service) String() string {
	return bs.name
}
