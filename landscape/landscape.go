package landscape

import (
	logging "github.com/ipfs/go-log"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"

	"github.com/mosaicdao/go-mosaic/threads"
)

var (
	log = logging.Logger("mosaic")
)

// Landscape provides a place for the node to explore and interact with
// In particular the landscape consists of contracts on (ethereum) chain(s)
// Specifically Peers are assigned to Boards on the contract
type Landscape interface {
	Peers(board threads.BoardID) []p2ppeer.AddrInfo
}

type defaultLandscape struct{}

func NewDefaultLandscape() Landscape {
	return &defaultLandscape{}
}

func (*defaultLandscape) Peers(threads.BoardID) []p2ppeer.AddrInfo {
	log.Panicf("Peers() not implemented for default landscape")
	return nil
}
