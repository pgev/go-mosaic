package reactor

import (
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type Reactor interface {
	// Service can be started and stopped
	service.Service
	MemberReactor
	PastUserReactor

	SetSwitch(*Switch)

	GetChannels() []*ChannelDescriptor
}

type MemberReactor interface {
	InitMember(member Member) Member

	AddMember(member Member)
}

type PastUserReactor interface {
	InitPastUser(pastUser PastUser) PastUser

	AddPastUser(pastUser PastUser)
}
