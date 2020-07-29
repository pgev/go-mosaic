package node

import (
	logging "github.com/ipfs/go-log"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/libs/service"
	thr "github.com/mosaicdao/go-mosaic/threads"
)

var (
	log = logging.Logger("node")
)

// Node

type Node struct {
	service.BaseService

	config  *cfg.Config

	threads thr.ThreadsNetwork
}

func NewNode(config *cfg.Config) (*Node, error) {
	// create DB etc

	threads := thr.NewThreadsNetwork(config.Threads)

	node := &Node{
		config:  config,
		threads: threads,
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
