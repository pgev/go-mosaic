package node

import (
	"github.com/ipfs/go-log/logging"
	cfg "github.com/mosaicdao/go-mosaic/config"
)

// Node Provider

type NodeProvider func(*cfg.Config, logging.Logger) (*Node, error)

func DefaultNewNode(config *cfg.Config, logger logging.Logger) (*Node, error) {

	return NewNode(config, logger)
}
