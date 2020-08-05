package explorer

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type theGraphExplorer struct {
	service.BaseService

	endpoint string
}

func NewTheGraphExplorer() Explorer {
	tge := &theGraphExplorer{}
	tge.BaseService = *service.NewBaseService("TheGraph Explorer", tge)
	return tge
}

func (*theGraphExplorer) OnStart() error {
	return nil
}

func (*theGraphExplorer) OnStop() {
}
