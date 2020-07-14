package node

import (
	"github.com/ipfs/go-log/logging"
	cfg "github.com/mosaicdao/go-mosaic/config"
)

// Node

type Node struct {
	// put DB
}

func NewNode(config *cfg.Config, logger logging.Logger) (*Node, error) {
	// create DB etc

	node := &Node{}
	return node, nil
}
