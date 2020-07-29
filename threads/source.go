package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Source interface {
	service.Service

	InitReactors(map[string]Reactor)

	Board() Board
	Sender() Sender
}
