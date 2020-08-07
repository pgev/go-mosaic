package landscape

import (
	"context"

	logging "github.com/ipfs/go-log"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"

	"github.com/mosaicdao/go-mosaic/libs/service"
	"github.com/mosaicdao/go-mosaic/threads"
)

var (
	log = logging.Logger("mosaic")
)

// Landscape provides a place for the node to explore and interact with
// In particular the landscape consists of contracts on (ethereum) chain(s)
// Specifically Peers are assigned to Boards on the contract
type Landscape interface {
	service.Service

	GetAssignments(p2ppeer.ID) ([]threads.BoardID, error)
	GetSources(threads.BoardID) []p2ppeer.ID
	GetPeers(threads.BoardID) []p2ppeer.AddrInfo
	Subscribe(ctx context.Context,
		options ...SubscriptionOption) <-chan LandscapeEvent
}
