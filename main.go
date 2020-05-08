package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	util "github.com/textileio/go-threads/util"
)

var (
	ctx context.Context
	ds  datastore.Batching
	net common.NetBoostrapper

	grey  = color.New(color.FgHiBlack).SprintFunc()
	green = color.New(color.FgHiGreen).SprintFunc()
	cyan  = color.New(color.FgHiCyan).SprintFunc()
	pink  = color.New(color.FgHiMagenta).SprintFunc()
	red   = color.New(color.FgHiRed).SprintFunc()
)

type notifee struct{}

func (n *notifee) HandlePeerFound(p peer.AddrInfo) {
	net.Host().Peerstore().AddAddrs(p.ID, p.Addrs, pstore.ConnectedAddrTTL)
}

func main() {
	repo := flag.String("repo", ".threads", "repo location")
	hostAddrStr := flag.String("hostAddr", "/ip4/0.0.0.0/tcp/4006", "Threads host bind address")
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.parse()

	hostAddr, err := ma.NewMultiAddress(*hostAddrStr)
	if err != nil {
		log.Fatal(err)
	}

	util.SetupDefaultLoggingConfig(*repo)
	if *debug {
		if err := logging.SetLogLevel("mosaic", "debug"); err != nil {
			log.Fatal(err)
		}
	}

	mosaicPath := filepath.Join(*repo, "mosaic")
	if err = os.MakedirAll(mosaicPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ds, err = ipfslite.BadgerDatastore(mosaicPath)
	if err != nil {
		log.Fatal(err)
	}

	net, err = common.DefaultNetwork(
		*repo,
		common.WithNetHostAdrr(hostAddr),
		common.WithNetDebug(*debug))
	if err != nil {
		log.Fatal(err)
	}
	defer net.Close()
	net.Bootstrap(util.DefaultBootstrapPeers())

	// Build a multicast DNS service
	ctx = context.Background()
	mdns, err := discovery.NewMdnsService(ctx, net.Host(), time.Second, "")
	if err != nil {
		log.Fatal(err)
	}
	defer mdns.Close()
	mdns.RegisterNotifee(&notifee{})

	// Start the prompt
	fmt.Println(grey("Welcome to Mosaic experiment 1: chat app with replication rings!!!"))
	fmt.Println(grey("Your peer ID is ") + green(net.Host().ID().String()))

	sub, err := net.Subscribe(ctx)
}
