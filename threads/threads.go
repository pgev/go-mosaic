package threads

import (
	"context"
	"time"

	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	badger "github.com/ipfs/go-ds-badger"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	cm "github.com/libp2p/go-libp2p-connmgr"
	corecm "github.com/libp2p/go-libp2p-core/connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-peerstore/pstoreds"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-threads/core/app"
	"github.com/textileio/go-threads/core/logstore"
	"github.com/textileio/go-threads/logstore/lstoreds"
	"github.com/textileio/go-threads/net"
	"google.golang.org/grpc"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

var (
	log = logging.Logger("threads")
)

type ThreadsNetwork interface {
	service.Service
}

// Threads provides a Threads network, and implements Servicable and Service.
type threads struct {
	service.BaseService

	cancel context.CancelFunc

	api app.Net // provides Connection

	config        *cfg.ThreadsConfig
	hostAddress   ma.Multiaddr
	peer          *ipfslite.Peer
	host          host.Host
	dht           *dual.DHT
	peerstore     peerstore.Peerstore
	litedatastore datastore.Datastore
	logdatastore  datastore.Datastore
}

// NewThreadsNetwork provides a ThreadsNetwork interface to a new instance.
// The ThreadsNetwork must be started by calling Start() before use.
func NewThreadsNetwork(privateNetworkKey crypto.PrivKey,
	config *cfg.ThreadsConfig) (ThreadsNetwork, error) {
	// TODO: take config for OnStart to work

	connManager := setupConnectionManager(
		config.ConnectionsLowWaterMark,
		config.ConnectionsHighWaterMark,
		config.ConnectionsGracePeriod,
	)

	hostAddress, err :=  config.HostAddress()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	// create IPFS Lite Peer
	litedatastore, peerstore, peer, host, dht, err := createIpfsLitePeer(
		ctx,
		config.IpfsLitePath(),
		privateNetworkKey,
		[]ma.Multiaddr{hostAddress},
		connManager,
	)
	if err != nil {
		cancel()
		return nil, err
	}

	// build a log store
	logdatastore, logstore, err := createLogStore(
		ctx,
		config.LogStorePath(),
	)
	if err != nil {
		cancel()
		litedatastore.Close()
		return nil, err
	}

	api, err := createNetworkAPI(
		ctx, host, peer, logstore,
		&net.Config{
			Debug: true,
		},
	)
	if err != nil {
		cancel()
		litedatastore.Close()
		logdatastore.Close()
		return nil, err
	}

	t := &threads{
		cancel:        cancel,
		api:           api,
		hostAddress:   hostAddress,
		peer:          peer,
		host:          host,
		dht:           dht,
		peerstore:     peerstore,
		litedatastore: litedatastore,
		logdatastore:  logdatastore,
	}
	return t, nil
}

func (t *threads) OnStart() error {
	// TODO: start ipfslite, logstore etc

	// subscribe to network: in a go routine
	// place to handle subscription updates (later)
	// incoming records of (new) logIds, parse messages and passed to switch
	return nil
}

func (t *threads) OnStop() {
	// close stuff

}

//------------------------------------------------------------------------------
// Private functions

func setupConnectionManager(low, high int,grace time.Duration) (
	corecm.ConnManager,
) {
	return cm.NewConnManager(low, high, grace)
}

func createIpfsLitePeer(
	ctx context.Context,
	ipfsLitePath string,
	privateNetworkKey crypto.PrivKey,
	listenAddresses []ma.Multiaddr,
	connectionManager corecm.ConnManager,
) (
	datastore.Datastore,
	peerstore.Peerstore,
	*ipfslite.Peer,
	host.Host,
	*dual.DHT,
	error,
) {
	// create litedatastore
	litedatastore, err := ipfslite.BadgerDatastore(ipfsLitePath)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	// create peerstore
	peerstore, err := pstoreds.NewPeerstore(ctx, litedatastore,
		 pstoreds.DefaultOpts())
	if err != nil {
		litedatastore.Close()
		return nil, nil, nil, nil, nil, err
	}
	host, dht, err := ipfslite.SetupLibp2p(
		ctx,
		privateNetworkKey,
		nil, // use an open libp2p network, let secret be nil
		listenAddresses,
		litedatastore,
		libp2p.Peerstore(peerstore),
		libp2p.ConnectionManager(connectionManager),
		libp2p.DisableRelay(),
	)
	if err != nil {
		litedatastore.Close()
		return nil, nil, nil, nil, nil, err
	}

	peer, err := ipfslite.New(ctx, litedatastore, host, dht, nil)
	if err != nil {
		litedatastore.Close()
		return nil, nil, nil, nil, nil, err
	}

	return litedatastore, peerstore, peer, host, dht, nil
}

func createLogStore(
	ctx context.Context,
	logStorePath string,
) (
	datastore.Datastore,
	logstore.Logstore,
	error,
) {
	logdatastore, err := badger.NewDatastore(logStorePath, &badger.DefaultOptions)
	if err != nil {
		return nil, nil, err
	}

	logstore, err := lstoreds.NewLogstore(ctx, logdatastore, lstoreds.DefaultOpts())
	if err != nil {
		logdatastore.Close()
		return nil, nil, err
	}

	return logdatastore, logstore, nil
}

func createNetworkAPI(
	ctx         context.Context,
	host        host.Host,
	peer        *ipfslite.Peer,
	logstore    logstore.Logstore,
	netConfig   *net.Config,
	grpcOptions ...grpc.ServerOption,
) (
	app.Net,
	error,
) {
	api, err := net.NewNetwork(
		ctx,
		host,
		peer.BlockStore(),
		peer,
		logstore,
		*netConfig,
		grpcOptions...
	)
	if err != nil {
		return nil, err
	}
	return api, nil
}
