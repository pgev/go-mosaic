package boards

import (
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
)

type Source struct {
	BoardID     BoardID
	Participant *p2ppeer.ID
}
