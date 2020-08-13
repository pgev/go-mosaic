package landscape

import (
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"

	"github.com/mosaicdao/go-mosaic/boards"
)

const (
	sourceChangeEventName   = "source_change_event"
	peerInfoUpdateEventName = "peer_addrinfo_update_event"
)

type LandscapeEvent interface {
	String() string
	BoardID() boards.BoardID
}

//------------------------------------------------------------------------------
// SourceChange event

type SourceChangeEvent struct {
	boardID boards.BoardID
}

func NewSourceChangeEvent(boardID boards.BoardID) LandscapeEvent {
	s := &SourceChangeEvent{
		boardID: boardID,
	}

	return s
}

func (*SourceChangeEvent) String() string {
	return sourceChangeEventName
}

func (s *SourceChangeEvent) BoardID() boards.BoardID {
	return s.boardID
}

//------------------------------------------------------------------------------
// PeerAddrInfoUpdate event

type PeerInfoUpdateEvent struct {
	boardID  boards.BoardID
	addrInfo p2ppeer.AddrInfo
	// distance between the location of the peer and the board
	// on the circuit
	distance uint64
}

func NewPeerInfoUpdateEvent(
	boardID boards.BoardID,
	addrInfo p2ppeer.AddrInfo,
	distance uint64,
) LandscapeEvent {
	return &PeerInfoUpdateEvent{
		boardID:  boardID,
		addrInfo: addrInfo,
		distance: distance,
	}
}

func (p *PeerInfoUpdateEvent) BoardID() boards.BoardID {
	return p.boardID
}

func (*PeerInfoUpdateEvent) String() string {
	return peerInfoUpdateEventName
}

func (p *PeerInfoUpdateEvent) Distance() uint64 {
	return p.distance
}
