package node

import (
	cfg "github.com/mosaicdao/go-mosaic/config"
)

// Node Provider

type NodeProvider func(*cfg.Config) (*Node, error)

func DefaultNewNode(config *cfg.Config) (*Node, error) {

	return NewNode(config)
}
