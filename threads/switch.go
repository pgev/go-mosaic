package threads

import (
	"fmt"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

// Switch defines a switch struct.
type Switch struct {
	service.BaseService

	reactors           map[string]Reactor
	channelDescriptors []*ChannelDescriptor
	reactorsByCh       map[byte]Reactor
}

func (sw *Switch) OnStart() error {
	// starts reactors
	for _, reactor := range sw.reactors {
		err := reactor.Start()
		if err != nil {
			return fmt.Errorf("failed to start %v: %w", reactor, err)
		}
	}

	return nil
}

func (sw *Switch) OnStop() {
	// stops reactors
	for _, reactor := range sw.reactors {
		reactor.Stop()
	}
}

// func (sw *Switch) AddReactor(name string, reactor Reactor) Reactor
// func (sw *Switch) RemoveReactor(name string, reactor Reactor)
// func (sw *Switch) Reactors() map[string]Reactor
// func (sw *Switch) Reactor(name string) Reactor
