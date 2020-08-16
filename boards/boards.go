package boards

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	badger "github.com/ipfs/go-ds-badger"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	cm "github.com/libp2p/go-libp2p-connmgr"
	corecm "github.com/libp2p/go-libp2p-core/connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	p2ppeer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/libp2p/go-libp2p-peerstore/pstoreds"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	mfma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-threads/core/app"
	"github.com/textileio/go-threads/core/logstore"
	"github.com/textileio/go-threads/logstore/lstoreds"
	"github.com/textileio/go-threads/net"
	"google.golang.org/grpc"

	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/libs/service"
)

var (
	grey  = color.New(color.FgHiBlack).SprintFunc()
	green = color.New(color.FgHiGreen).SprintFunc()

	log = logging.Logger("boards")
)

type notifee struct {
	b BoardsManager
}

func (n *notifee) HandlePeerFound(p p2ppeer.AddrInfo) {
	log.Infof("found peer %v, adding to peerstore", p)
	n.b.Host().Peerstore().AddAddrs(
		p.ID,
		p.Addrs,
		peerstore.ConnectedAddrTTL,
	)
}

type BoardsManager interface {
	service.Service

	Host() host.Host
	Peerstore() peerstore.Peerstore
}

// Boards implements a BoardsManager, and implements Servicable and Service.
type boards struct {
	service.BaseService

	// cancel dependent goroutines
	childCancel context.CancelFunc

	api app.Net // provides threads network connection to boards manager

	hostAddress    mfma.Multiaddr
	bootstrapPeers []p2ppeer.AddrInfo
	peer           *ipfslite.Peer
	host           host.Host
	dht            *dual.DHT
	peerstore      peerstore.Peerstore
	liteDatastore  datastore.Datastore
	logDatastore   datastore.Datastore
	viewDatastore  datastore.Datastore

	mdns discovery.Service
}

var _ (BoardsManager) = (*boards)(nil)

// NewBoardsManager provides a BoardsManager interface to a new instance of boards.
// The BoardsManager must be started by calling Start() before use.
func NewBoardsManager(
	ctx context.Context,
	privateNetworkKey crypto.PrivKey,
	bootstrapPeers []p2ppeer.AddrInfo,
	config *cfg.ThreadsConfig) (BoardsManager, error) {
	// TODO: take config for OnStart to work

	connManager := setupConnectionManager(
		config.ConnectionsLowWaterMark,
		config.ConnectionsHighWaterMark,
		config.ConnectionsGracePeriod,
	)

	hostAddress, err := config.HostAddress()
	if err != nil {
		return nil, err
	}

	childCtx, childCancel := context.WithCancel(ctx)

	_, err = createViewDatastore(
		childCtx,
		config.ViewStorePath(),
	)

	// create IPFS Lite Peer
	liteDatastore, peerstore, peer, host, dht, err := createIpfsLitePeer(
		childCtx,
		config.IpfsLitePath(),
		privateNetworkKey,
		[]mfma.Multiaddr{hostAddress},
		connManager,
	)
	if err != nil {
		childCancel()
		return nil, err
	}

	// build a log store
	logDatastore, logstore, err := createLogStore(
		childCtx,
		config.LogStorePath(),
	)
	if err != nil {
		childCancel()
		liteDatastore.Close()
		return nil, err
	}

	api, err := createNetworkAPI(
		childCtx, host, peer, logstore,
		&net.Config{
			Debug: true,
		},
	)
	if err != nil {
		childCancel()
		liteDatastore.Close()
		logDatastore.Close()
		return nil, err
	}

	b := &boards{
		childCancel:    childCancel,
		api:            api,
		hostAddress:    hostAddress,
		bootstrapPeers: bootstrapPeers,
		peer:           peer,
		host:           host,
		dht:            dht,
		peerstore:      peerstore,
		liteDatastore:  liteDatastore,
		logDatastore:   logDatastore,
	}
	b.BaseService = *service.NewBaseService("BoardsManager", b)

	go b.autoclose(ctx)

	return b, nil
}

func (b *boards) OnStart() error {

	// bootstrap peers; for bigfish project, piggyback on the threadsDB
	// and IPFS public bootstrap peers
	// TODO: refine where to bootstrap from depending on known Column members
	b.peer.Bootstrap(b.bootstrapPeers)

	// Build a MDNS service
	ctx := context.Background()
	mdns, err := discovery.NewMdnsService(ctx, b.api.Host(), time.Second, "")
	if err != nil {
		log.Warnf("fatal error creating MDNS service: %w", err)
		return err
	}
	notifee := &notifee{
		b: b,
	}
	mdns.RegisterNotifee(notifee)
	b.mdns = mdns
	// Start the prompt
	fmt.Println(grey("Welcome to MOSAIC!"))
	fmt.Println(grey("Your peer ID is ") + green(b.Host().ID().String()))
	fmt.Printf("Listening on addresses: %v\n", b.Host().Addrs())

	// subscribe to threads network without any filter options
	sub, err := b.api.Subscribe(ctx)
	if err != nil {
		log.Errorf("failed to subscribe to threads network: %w", err)
	}

	go func() {
		for rec := range sub {
			fmt.Printf("got new record: %v", rec)
		}
	}()

	// place to handle subscription updates (later)
	// incoming records of (new) logIds, parse messages and passed to switch
	return nil
}

func (b *boards) OnStop() {
	// close all depending goroutines
	b.close()
}

func (b *boards) Host() host.Host {
	return b.host
}

func (b *boards) Peerstore() peerstore.Peerstore {
	return b.peerstore
}

//------------------------------------------------------------------------------
// Private functions

func setupConnectionManager(low, high int, grace time.Duration) corecm.ConnManager {
	return cm.NewConnManager(low, high, grace)
}

func createViewDatastore(
	ctx context.Context,
	viewDatastorePath string,
) (
	datastore.Datastore,
	error,
) {
	vs, err := badger.NewDatastore(viewDatastorePath, &badger.DefaultOptions)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func createIpfsLitePeer(
	ctx context.Context,
	ipfsLitePath string,
	privateNetworkKey crypto.PrivKey,
	listenAddresses []mfma.Multiaddr,
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
	ctx context.Context,
	host host.Host,
	peer *ipfslite.Peer,
	logstore logstore.Logstore,
	netConfig *net.Config,
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
		grpcOptions...,
	)
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (b *boards) autoclose(ctx context.Context) {
	<-ctx.Done()
	log.Info("threads network autoclosing")
	b.close()
}

func (b *boards) close() {
	// close datastores and dependencies
	if err := b.api.Close(); err != nil {
		log.Warnf("error closing threads network API: %w", err)
	}
	// IPFSLite t.peer is autoclosed by cancel()
	b.childCancel()
	if err := b.dht.Close(); err != nil {
		log.Warnf("error closing DHT: %w", err)
	}
	if err := b.host.Close(); err != nil {
		log.Warnf("errror closing host: %w", err)
	}
	if err := b.peerstore.Close(); err != nil {
		log.Warnf("error closing peerstore: %w", err)
	}
	if err := b.liteDatastore.Close(); err != nil {
		log.Warnf("error closing litedatastore: %w", err)
	}
	if err := b.logDatastore.Close(); err != nil {
		log.Warnf("error closing logdatastore: %w", err)
	}
	if err := b.mdns.Close(); err != nil {
		log.Warnf("error closing mdns: %w", err)
	}
}
