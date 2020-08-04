package node

import (
	"context"

	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	txtutil "github.com/textileio/go-threads/util"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/landscape"
	sgn "github.com/mosaicdao/go-mosaic/signer"
)

// Node Provider

type NodeProvider func(context.Context, *cfg.Config) (*Node, landscape.Landscape, error)

func DefaultNewNode(ctx context.Context, config *cfg.Config) (*Node, landscape.Landscape, error) {
	// TODO: replace nils

	l := landscape.NewDefaultLandscape()


	n, err := NewNode(ctx, defaultNetworkPrivateKey(config), nil, config)
	if err != nil {
		log.Panicf("failed to create a new node: %w", err)
	}

	return n, l, err
}

func defaultNetworkPrivateKey(config *cfg.Config) NetworkPrivateKeyProvider {
	return func(config *cfg.Config) (p2pcrypto.PrivKey, error) {
		// Generate new random key for libp2p; TODO: load key from disk
		netKey, err := sgn.GenerateFileNetworkSigner(config.NodePrivateKeyFile())
		if err != nil {
			return nil, err
		}
		return netKey, nil
	}
}

func defaultBootstrapPeers() BootstrapListProvider {
	return func() []p2ppeer.AddrInfo {
		return txtutil.DefaultBoostrapPeers()
	}
}
