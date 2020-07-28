package threads

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type BaseReactor struct {
	service.BaseService

	Switch *Switch
}

func NewBaseReactor(name string, impl service.Servicable) *BaseReactor {
	return &BaseReactor{
		BaseService: *service.NewBaseService(name, impl),
		Switch:      nil,
	}
}

func (r *BaseReactor) SetSwitch(sw *Switch) { r.Switch = sw }

func (*BaseReactor) GetTopics() []Topic { return nil }

func (*BaseReactor) AddDatabus(databus Databus)                                {}
func (*BaseReactor) RemoveDatabus(databus Databus)                             {}
func (*BaseReactor) ReceiveFrom(databus Databus, topic Topic, msgBytes []byte) {}
func (*BaseReactor) InitDatabus(databus Databus) Databus                       { return databus }
