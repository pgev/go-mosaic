package node

import (
	"context"
	"os"
	"strconv"

	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/landscape"
)

// Cut Corners Node Provider cuts corners
// it uses hardcoded keys and IP addresses to bootstrap a local network
// and help us test the initial wiring while building the code base

// CutCornersNewNode provides a new node, but override bootstrap and private key
// functionality
func CutCornersNewNode(ctx context.Context, config *cfg.Config) (*Node, landscape.Landscape, error) {

	ccl := landscape.CreateCutCornersLandscape()

	n, err := NewNode(
		ctx, cutCornerNetworkPrivateKey(ccl), cutCornerBootstrapPeers(ccl), config,
	)
	if err != nil {
		log.Panicf("failed to create a new node: %w", err)
	}
	// read env and set to private var
	return n, ccl, nil
}

func cutCornerNetworkPrivateKey(ccl *landscape.CutCornersLandscape) NetworkPrivateKeyProvider {
	return func(*cfg.Config) (p2pcrypto.PrivKey, error) {
		is := os.Getenv("MOSAIC_CC_ID")
		index64, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			log.Panicf("failed to get index from env variable MOSAIC_CC_ID (%s): %w", is, err)
		}
		index := int(index64)
		privKey := ccl.GetPrivateKey(index)
		return privKey, nil
	}
}

func cutCornerBootstrapPeers(ccl *landscape.CutCornersLandscape) BootstrapListProvider {
	return func() []p2ppeer.AddrInfo {
		return ccl.GetBootstrapPeers()
	}
}
