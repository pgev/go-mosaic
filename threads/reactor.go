package reactor

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type MemberReactor interface {
	AddMember(member Member)
	RemoveMember(member Member)
	ReceiveFromMember(chID byte, member Member, msgBytes []byte)
	InitMember(member Member) Member
}

type PastUserReactor interface {
	AddPastUser(pastUser PastUser)
	RemovePastUser(pastUser PastUser)
	ReceiveFromPastUser(chID byte, pastUser PastUser, msgBytes []byte)
	InitPastUser(pastUser PastUser) PastUser
}

type Reactor interface {
	service.Servicable

	MemberReactor
	PastUserReactor

	SetSwitch(*Switch)

	GetChannels() []*ChannelDescriptor
}

type BaseReactor struct {
	service.BaseService

	Switch *Switch
}

func NewBaseReactor(name string, impl Reactor) *BaseReactor {
	return &BaseReactor{
		BaseService: *service.NewBaseService(name, impl),
		Switch: 	 nil
	}
}

func (r *BaseReactor) SetSwitch (sw *Switch) { r.Switch = sw }

func (*BaseReactor) GetChannels() []*ChannelDescriptor 								{ return nil }

func (*BaseReactor) AddMember(member Member) 										{}
func (*BaseReactor) RemoveMember(member Member) 									{}
func (*BaseReactor) ReceiveFromMember(chID byte, member Member, msgBytes []byte) 	{}
func (*BaseReactor) InitMember(member Member) Member								{ return member}

func (*BaseReactor) AddPastUser(pastUser PastUser) 										{}
func (*BaseReactor) RemovePastUser(pastUser PastUser) 									{}
func (*BaseReactor) ReceiveFromPastUser(chID byte, pastUser PastUser, msgBytes []byte) 	{}
func (*BaseReactor) InitPastUser(pastUser PastUser) PastUser							{ return pastUser}
