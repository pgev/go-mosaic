package node

import (
	logging "github.com/ipfs/go-log"

	cfg "github.com/mosaicdao/go-mosaic/config"
)

var (
	log = logging.Logger("node")
)

// Node

type Node struct {
	// put DB

}

func NewNode(config *cfg.Config) (*Node, error) {
	// create DB etc

	node := &Node{}
	return node, nil
}
