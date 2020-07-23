package threads

import (
	"github.com/mosaicdao/go-mosaic/column"
	"github.com/mosaicdao/go-mosaic/gate"
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

func (*BaseReactor) GetChannels() []*ChannelDescriptor { return nil }

func (*BaseReactor) AddMember(member column.Member)                                     {}
func (*BaseReactor) RemoveMember(member column.Member)                                  {}
func (*BaseReactor) ReceiveFromMember(chID byte, member column.Member, msgBytes []byte) {}
func (*BaseReactor) InitMember(member column.Member) column.Member                      { return member }

func (*BaseReactor) AddPastUser(pastUser gate.PastUser)                                     {}
func (*BaseReactor) RemovePastUser(pastUser gate.PastUser)                                  {}
func (*BaseReactor) ReceiveFromPastUser(chID byte, pastUser gate.PastUser, msgBytes []byte) {}
func (*BaseReactor) InitPastUser(pastUser gate.PastUser) gate.PastUser                      { return pastUser }
