package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Reactor interface {
	service.Service

	InitSource(source Source) (Source, error)
	AddSource(source Source) error

	ReceiveMsg(msg *Message)

	Board() *Board
	GetTopics() []*Topic

	SetSwitch(*Switch)
}
