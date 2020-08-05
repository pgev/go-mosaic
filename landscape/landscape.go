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
	Peers(board threads.BoardID) []p2ppeer.AddrInfo
	// Subscription to source changes of (one, or more or all) board(s)
	SubscribeSourceChange(ctx context.Context,
		options ...SubscriptionOption) (<-chan SourceChange, error)
	SubscribeLogAppend(ctx context.Context,
		options ...SubscriptionOption) (<-chan BoardLog, error)
}

type defaultLandscape struct {
	service.BaseService
	// TODO: landscape has datastore for board memberships
}

// NewDefaultLandscape provides a (currently empty!) default landscape
// TODO: default to Ethereum or testnet Goerli
func NewDefaultLandscape() Landscape {

	dl := &defaultLandscape{}
	dl.BaseService = *service.NewBaseService("DefaultLandscape", dl)

	return dl
}

func (*defaultLandscape) Peers(threads.BoardID) []p2ppeer.AddrInfo {
	log.Panicf("Peers() not implemented for default landscape")
	return nil
}

func (*defaultLandscape) OnStart() error {
	return nil
}

func (*defaultLandscape) OnStop() {}
