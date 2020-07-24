package threads

import (
	"github.com/mosaicdao/go-mosaic/column"
	"github.com/mosaicdao/go-mosaic/gate"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

type MemberReactor interface {
	AddMember(member column.Member)
	RemoveMember(member column.Member)
	InitMember(member column.Member) column.Member

	ReceiveFromMember(topicID TopicID, member column.Member, msgBytes []byte)
}

type PastUserReactor interface {
	AddPastUser(pastUser gate.PastUser)
	RemovePastUser(pastUser gate.PastUser)
	InitPastUser(pastUser gate.PastUser) gate.PastUser

	ReceiveFromPastUser(chID byte, pastUser gate.PastUser, msgBytes []byte)
}

type Reactor interface {
	service.Service

	MemberReactor
	PastUserReactor

	SetSwitch(*Switch)
	GetTopics() []*Topic
}
