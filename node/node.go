package node

import (
	"context"

	logging "github.com/ipfs/go-log"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"

	cfg "github.com/mosaicdao/go-mosaic/config"
	lnd "github.com/mosaicdao/go-mosaic/landscape"
	"github.com/mosaicdao/go-mosaic/libs/service"
	sc "github.com/mosaicdao/go-mosaic/scout"
	thr "github.com/mosaicdao/go-mosaic/threads"
)

var (
	log = logging.Logger("node")
)

// Node combines all the services running to operate as a member
type Node struct {
	service.BaseService

	config *cfg.Config

	childCancel context.CancelFunc

	// for bigfish, don't yet create a keystore, rather just keep the explicit
	// keys in the node. TODO: improve key management
	networkKey p2pcrypto.PrivKey

	threads thr.ThreadsNetwork
	scout   sc.Scout
}

type NetworkPrivateKeyProvider func(config *cfg.Config) (p2pcrypto.PrivKey, error)

type BootstrapListProvider func() []p2ppeer.AddrInfo

// NewNode creates a new node based on the configuration provided
func NewNode(ctx context.Context,
	landscape lnd.Landscape,
	networkPrivateKeyProvider NetworkPrivateKeyProvider,
	bootstrapListProvider BootstrapListProvider,
	config *cfg.Config) (*Node, error) {
	childCtx, childCancel := context.WithCancel(ctx)

	netKey, err := networkPrivateKeyProvider(config)
	if err != nil {
		// child context is not yet used, but for clarity clean it up
		childCancel()
		return nil, err
	}

	threads, err := thr.NewThreadsNetwork(
		childCtx, netKey, bootstrapListProvider(), config.Threads,
	)
	if err != nil {
		// threads creation failed, so it has cleaned up,
		// but for clarity clean it up
		childCancel()
		return nil, err
	}

	scout, err := sc.NewScout(childCtx, landscape, threads)
	if err != nil {
		childCancel()
		return nil, err
	}

	node := &Node{
		config:      config,
		childCancel: childCancel,
		networkKey:  netKey,
		threads:     threads,
		scout:       scout,
	}
	node.BaseService = *service.NewBaseService("Node", node)

	go node.autoclose(ctx)

	return node, nil
}

func (n *Node) OnStart() error {
	if err := n.threads.Start(); err != nil {
		// TODO: properly close off threads which already has datastores active
		log.Errorf("failed to start threads: %w", err)
		return err
	}

	if err := n.scout.Start(); err != nil {
		log.Errorf("failed to start scout: %w", err)
		return err
	}
	return nil
}

func (n *Node) OnStop() {
	n.close()
}

func (n *Node) Threads() thr.ThreadsNetwork {
	return n.threads
}

//------------------------------------------------------------------------------
// Private functions

func (n *Node) autoclose(ctx context.Context) {
	<-ctx.Done()
	n.close()
}

func (n *Node) close() {
	n.threads.Stop()
	n.childCancel()
}
