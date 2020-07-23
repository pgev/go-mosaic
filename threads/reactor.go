package threads

import (
	"github.com/mosaicdao/go-mosaic/column"
	"github.com/mosaicdao/go-mosaic/gate"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type MemberReactor interface {
	AddMember(member column.Member)
	RemoveMember(member column.Member)
	ReceiveFromMember(chID byte, member column.Member, msgBytes []byte)
	InitMember(member column.Member) column.Member
}

type PastUserReactor interface {
	AddPastUser(pastUser gate.PastUser)
	RemovePastUser(pastUser gate.PastUser)
	ReceiveFromPastUser(chID byte, pastUser gate.PastUser, msgBytes []byte)
	InitPastUser(pastUser gate.PastUser) gate.PastUser
}

type Reactor interface {
	service.Service

	MemberReactor
	PastUserReactor

	SetSwitch(*Switch)
	GetChannels() []*ChannelDescriptor
}
