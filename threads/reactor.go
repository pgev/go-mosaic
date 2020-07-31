package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Reactor interface {
	service.Service

	InitSource(source *Source) error
	AddSource(source *Source) error

	ReceiveMsg(msg *Message)

	BoardId() BoardId
	GetTopicIds() []TopicId

	SetSwitch(*Switch)
}
