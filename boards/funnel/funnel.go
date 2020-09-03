package funnel

import (
	"github.com/mosaicdao/go-mosaic/boards"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Funnel struct {
	service.BaseService

	boardID boards.BoardID

}

func (f *Funnel) OnStart() error {
	// start listening to the threads api
}
