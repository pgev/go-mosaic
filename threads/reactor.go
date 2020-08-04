package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Reactor interface {
	service.Service

	InitSource(source *Source) error
	AddSource(source *Source) error

	ReceiveMsg(msg *Message)

	BoardID() BoardID
	GetTopicIDs() []TopicID

	SetSwitch(*Switch)
}
