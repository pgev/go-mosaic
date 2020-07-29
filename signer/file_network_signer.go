package signer

import (
	"errors"
	"crypto/rand"

	"github.com/libp2p/go-libp2p-core/crypto"
	pb "github.com/libp2p/go-libp2p-core/crypto/pb"
)

// FileNetworkSigner implements libp2p crypto.PrivKey such that it can sign
// for the p2p and logs. It panics however when attempting to access the private
// bytes over Bytes(), Raw() or Equal().
// FileNetworkSigner is always of key type ed25519 because go-threads
// explicitly uses ed25519 for signing JWT tokens
type FileNetworkSigner struct {
	k crypto.PrivKey

	path string
}

var (
	// ErrPrivateKeyBytesAccessDenied is returned when private key bytes are accessed
	ErrPrivateKeyBytesAccessDenied = errors.New("Access to private key bytes is not supported")
)

var _ (crypto.PrivKey) = (*FileNetworkSigner)(nil)
var _ (Signer) = (*FileNetworkSigner)(nil)

// NewNetworkSignerFromFile returns a NetworkSigner from a private key stored
// on disk, and allows the user to access the private key
func NewNetworkSignerFromFile(path string) (*FileNetworkSigner, error) {
	return &FileNetworkSigner{}, nil
}

// GenerateFileNetworkSigner generates a random ed25519 private key and
// returns it as a FileNetworkSigner. It stores the private key to the file path.
// It returns an error if the file already exists.
func GenerateFileNetworkSigner(path string) (*FileNetworkSigner, error) {
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}
	publicBytes, _ := priv.GetPublic().Bytes()
	log.Infof("public network key generated: 0x%x", publicBytes)
	// TODO: store to path
	networkSigner := &FileNetworkSigner{
		k:    priv,
		path: path,
	}
	return networkSigner, nil
}

// Type of the private key is ed25519
func (ns *FileNetworkSigner) Type() pb.KeyType {
	return ns.k.Type()
}

// Bytes will return an error, blocking access to the private key bytes
func (ns *FileNetworkSigner) Bytes() ([]byte, error) {
	return ns.denyRawBytesAccess()
}

// Raw will return an error, blocking access to the private key bytes
func (ns *FileNetworkSigner) Raw() ([]byte, error) {
	return ns.denyRawBytesAccess()
}

// Equals will return an error, because comparison (now) happens on private key
// and access is restricted
func (ns *FileNetworkSigner) Equals(o crypto.Key) bool {
	log.Panic("attempt to access Equal with impl using private key of %v", ns.k.GetPublic())
	return false
}

// Sign returns a signature from an input message.
func (ns *FileNetworkSigner) Sign(msg []byte) ([]byte, error) {
	return ns.k.Sign(msg)
}

// GetPublic returns an ed25519 public key from the private key.
func (ns *FileNetworkSigner) GetPublic() crypto.PubKey {
	return ns.k.GetPublic()
}

//------------------------------------------------------------------------------
// Private functions

func (ns *FileNetworkSigner) denyRawBytesAccess() ([]byte, error) {
	// while in development, actually panic to make sure no libraries
	// try to access the bytes of the private key directly
	log.Panic("attempt to access bytes of private key of %v", ns.k.GetPublic())
	return nil, ErrPrivateKeyBytesAccessDenied
}
