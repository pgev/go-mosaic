package threads

import (
	"github.com/mosaicdao/go-mosaic/column"
	"github.com/mosaicdao/go-mosaic/gate"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

// MemberReactor defines an interface of a member reactor.
type MemberReactor interface {
	AddMember(member column.Member)
	RemoveMember(member column.Member)
	InitMember(member column.Member) column.Member

	ReceiveFromMember(topicID TopicID, member column.Member, msgBytes []byte)
}

// PastUserReactor defines an interface of a past user reactor.
type PastUserReactor interface {
	AddPastUser(pastUser gate.PastUser)
	RemovePastUser(pastUser gate.PastUser)
	InitPastUser(pastUser gate.PastUser) gate.PastUser

	ReceiveFromPastUser(topicID TopicID, pastUser gate.PastUser, msgBytes []byte)
}

// Reactor defines an interface of a reactor.
type Reactor interface {
	service.Service

	MemberReactor
	PastUserReactor

	GetTopics() []*Topic

	SetSwitch(*Switch)
}
