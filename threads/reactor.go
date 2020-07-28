package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Reactor interface {
	service.Service

	InitDatabus(databus Databus) Databus
	AddDatabus(databus Databus) error
	RemoveDatabus(databus Databus)
	ReceiveFrom(databus Databus, topic Topic, msgBytes []byte)

	GetTopics() []Topic
	SetSwitch(*Switch)
}
