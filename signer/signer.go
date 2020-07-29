package signer

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	logging "github.com/ipfs/go-log"
)

var (
	log = logging.Logger("signer")
)

// Signer provides an overlapping interface with a subset of libp2p
// crypto.PrivKey, Sign([]byte) and GetPublic(), but it excludes crypto.Key
// to better shield private keys.
type Signer interface {
	Sign(msg []byte) ([]byte, error)

	GetPublic() crypto.PubKey
}
