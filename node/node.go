package node

import (
	logging "github.com/ipfs/go-log"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

var (
	log = logging.Logger("node")
)

// Node

type Node struct {
	service.BaseService
	// put DB
}

func NewNode(config *cfg.Config) (*Node, error) {
	// create DB etc

	node := &Node{}
	node.BaseService = *service.NewBaseService("Node", node)

	return node, nil
}

func (n *Node) OnStart() error {
	return nil
}

func (n *Node) OnStop() {

}
