package node

import (
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

	config  *cfg.Config

	// for bigfish, don't yet create a keystore, rather just keep the explicit
	// keys in the node. TODO: improve key management
	networkKey *sgn.FileNetworkSigner

	threads thr.ThreadsNetwork
}

// NewNode creates a new node based on the configuration provided
func NewNode(config *cfg.Config) (*Node, error) {
	// create DB etc

	// Generate new random key for libp2p; TODO: load key from disk
	netKey, err := sgn.GenerateFileNetworkSigner(config.NodePrivateKeyFile())
	if err != nil {
		return nil, err
	}

	threads := thr.NewThreadsNetwork(netKey, config.Threads)

	node := &Node{
		config:     config,
		networkKey: netKey,
		threads:    threads,
	}
	node.BaseService = *service.NewBaseService("Node", node)

	return node, nil
}

func (n *Node) OnStart() error {
	// TODO: create threads network

	return nil
}

func (n *Node) OnStop() {

}
