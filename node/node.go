package node

import (
	"context"
	logging "github.com/ipfs/go-log"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/libs/service"
	sgn "github.com/mosaicdao/go-mosaic/signer"
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
	networkKey *sgn.FileNetworkSigner

	threads thr.ThreadsNetwork
}

// NewNode creates a new node based on the configuration provided
func NewNode(ctx context.Context, config *cfg.Config) (*Node, error) {
	childCtx, childCancel := context.WithCancel(ctx)

	// Generate new random key for libp2p; TODO: load key from disk
	netKey, err := sgn.GenerateFileNetworkSigner(config.NodePrivateKeyFile())
	if err != nil {
		return nil, err
	}

	threads, err := thr.NewThreadsNetwork(childCtx, netKey, config.Threads)
	if err != nil {
		return nil, err
	}

	node := &Node{
		config:      config,
		childCancel: childCancel,
		networkKey:  netKey,
		threads:     threads,
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
	return nil
}

func (n *Node) OnStop() {
	n.close()
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
