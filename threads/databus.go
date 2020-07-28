package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

// DatabusID defines a type of databus id.
type DatabusID []byte

type Databus interface {
	service.Service

	ID() DatabusID
	InitReactors(reactorsByTopicID map[TopicID]Reactor)
	compareByID(Databus) bool
}
