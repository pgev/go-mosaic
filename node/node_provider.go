package node

import (
	"context"

	cfg "github.com/mosaicdao/go-mosaic/config"
)

// Node Provider

type NodeProvider func(context.Context, *cfg.Config) (*Node, error)

func DefaultNewNode(ctx context.Context, config *cfg.Config) (*Node, error) {

	return NewNode(ctx, config)
}
