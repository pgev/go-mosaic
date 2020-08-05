package landscape

import (
	"encoding/hex"

	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	mfma "github.com/multiformats/go-multiaddr"
	txtthread "github.com/textileio/go-threads/core/thread"

	"github.com/mosaicdao/go-mosaic/threads"
)

// CutCornersLandscape is a temporary struct to have a hard-coded, shared threadsId
// and service and read key for the shared logs
type CutCornersLandscape struct {
	ThreadsID txtthread.ID
	Key       txtthread.Key
}

var (
	threadID = "bafkw3ulx7jrkuflaent7rrqkw7jaf4pdxhllvz4c32yn3mbusuw3pyi"

	threadKeys = "batgpncb7nve5skmo2y24nsccvoyh3ihkdazmkjbzm6zk7vwr7ykuhee3os5zyp6buihwv33fb2pwonjr4rddnrt4csle6ojgnfbhqya"

	memberPrivateKeyStrings = []string{
		"e4b404b19b59749a92141d1f1ef22509ac01480148923a1f7c9f65e68e80b85f97ea5ddec354513941796d00085c66daa866da65aa2a548a5fa2f0b7388823a4",
		"2c56a4518f3aa4add8f83d149a23a75a398018b59c3d01084626d83753189d2bdee10a66d305bb4d33d78dbb14a882b44ebe0ffd647d696c53c3181e7d20f4ca",
		// "42d18b2ac8a1a60a9bccb2995233ae3c46e15a595343ad445bd9165458c38cc86e846ac9b8dddd4831b79552b62d006246052238f7fb9e6ea8ea4a5d73767503",
		// "56e379081af7c1fb0820e78d45850d665f971c2e2bfc3df836ed8f074b9c4ad9fc155f8941bfa50e5d10f31e1d0212e9d952c09f505d25df6dcbaddca74c409b",
	}

	memberLocalIPAddrs = []string{
		"/ip4/0.0.0.0/tcp/4007",
		"/ip4/0.0.0.0/tcp/4008",
		// "/ip4/0.0.0.0/tcp/4009",
		// "/ip4/0.0.0.0/tcp/4010",
	}

	memberPublicKeys = []p2pcrypto.PubKey{}
	memberPeerIDs    = []p2ppeer.ID{}
	memberAddrInfos  = []p2ppeer.AddrInfo{}
)

func init() {
	var err error

	for _, privks := range memberPrivateKeyStrings {
		privk := unmarshalPrivateKey(privks)
		pubk := privk.GetPublic()
		peerID, err := p2ppeer.IDFromPublicKey(pubk)
		if err != nil {
			log.Panicf("failed to get PeerID from PublicKey: %w", err)
		}
		memberPublicKeys = append(memberPublicKeys, pubk)
		memberPeerIDs = append(memberPeerIDs, peerID)
	}

	var memberMultiaddrs []mfma.Multiaddr

	for index, ip := range memberLocalIPAddrs {
		ma, err := mfma.NewMultiaddr(ip + "/p2p/" + memberPeerIDs[index].String())
		if err != nil {
			log.Panicf("failed to create a multiaddress from a local ip address %v", ip)
		}
		memberMultiaddrs = append(memberMultiaddrs, ma)
	}

	// sort any duplicated PeerIds into the same address info
	memberAddrInfos, err = p2ppeer.AddrInfosFromP2pAddrs(memberMultiaddrs...)
	if err != nil {
		log.Panicf("failed to initialise member address info in cutcorner landscape: %w", err)
	}
}

// CreateCutCornersLandscape creates an new CutCorners Landscape with
// corners cut; ie. hardcoded private keys and service, read keys for a thread
// TODO: delete all this once components are wired together and basic testing
// is in place
func CreateCutCornersLandscape() *CutCornersLandscape {
	// set a hardcoded threadID (just a random number)
	id, err := txtthread.Decode(threadID)
	if err != nil {
		// this is just for crude scaffolding while building code
		panic(err)
	}

	// set a hardcoded service and read key, again just for bootstrapping
	// the codebase and making it easy to assert the lower parts of
	// the stack are wired together sensibly
	k, err := txtthread.KeyFromString(threadKeys)
	if err != nil {
		panic(err)
	}
	return &CutCornersLandscape{
		ThreadsID: id,
		Key:       k,
	}
}

// GetPrivateKey returns a hardcoded private key, cutting corners 👻
func (*CutCornersLandscape) GetPrivateKey(index int) p2pcrypto.PrivKey {
	if index < 0 || index > len(memberPrivateKeyStrings) {
		log.Panicf("Index (%v) is out of range", index)
	}

	return unmarshalPrivateKey(memberPrivateKeyStrings[index])
}

// GetBootstrapPeers returns the hardcoded member address info to use
// as bootstrap peers, because cutting corners 👻
func (*CutCornersLandscape) GetBootstrapPeers() []p2ppeer.AddrInfo {
	if len(memberAddrInfos) == 0 {
		log.Panicf("no bootstrap peers in cut corner landscape (%v)", memberAddrInfos)
	}
	// allpeers := append(txtutil.DefaultBoostrapPeers(), memberAddrInfos...)
	return memberAddrInfos
}

// Peers provides a set of hardcoded address info, cutting corners
func (*CutCornersLandscape) Peers(threads.BoardID) []p2ppeer.AddrInfo {
	return memberAddrInfos
}

//------------------------------------------------------------------------------
// Private functions

func unmarshalPrivateKey(k string) p2pcrypto.PrivKey {
	b, err := hex.DecodeString(k)
	if err != nil {
		log.Panicf("failed to decode hexstring to bytes: %w", err)
	}
	privk, err := p2pcrypto.UnmarshalEd25519PrivateKey(b)
	if err != nil {
		log.Panicf("failed to unmarshal bytes to private ed25519 key: %w", err)
	}
	return privk
}
