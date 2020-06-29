package main

import (
	"os"
	"time"

	logging "github.com/ipfs/go-log"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/namsral/flag"
	threadsCommon "github.com/textileio/go-threads/common"
	threadsUtil "github.com/textileio/go-threads/util"
)

var log = logging.Logger("mosaicd")

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "MOSAIC", 0)

	repo := fs.String("repo", ".mosaic", "Mosaic repository location")
	threadsRepo := fs.String("threadsRepo", ".mosaic/threads",
		"ThreadsDB repo location")
	hostAddrStr := fs.String("hostAddr", "/ip4/0.0.0.0/tcp/4006",
		"Threads host bind address")
	apiAddrStr := fs.String("apiAddr", "/ip4/127.0.0.1/tcp/6500",
		"Mosaic API bind address")
	threadsAPIAddrStr := fs.String("threadsAPIAddr",
		"/ip4/127.0.0.1/tcp/6006", "Threads API bind address")
	connectionsLowWater := fs.Int("connectionsLowWater", 100,
		"Low watermark of libp2p connections that will be maintained")
	connectionsHighWater := fs.Int("connectionsHighWater", 400,
		"High watermark of libp2p connections that will be maintained")
	connectionsGracePeriod := fs.Duration("connectionsGracePeriod", time.Second*20,
		"Duration a newly opened connection is given before it becomes subject to pruning")
	debug := fs.Bool("debug", false, "Enable debug logging")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	hostAddr, err := ma.NewMultiaddr(*hostAddrStr)
	if err != nil {
		log.Fatal(err)
	}
	// apiAddr, err := ma.NewMultiaddr(*apiAddrStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// threadsAPIAddr, err := ma.NewMultiaddr(*threadsAPIAddrStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	threadsUtil.SetupDefaultLoggingConfig(*threadsRepo)
	if *debug {
		if err := logging.SetLogLevel("mosaicd", "debug"); err != nil {
			log.Fatal(err)
		}
	}

	log.Debugf("Mosaic repository: %v", *repo)
	log.Debugf("Threads repository: %v", *threadsRepo)
	log.Debugf("Host address: %v", *hostAddrStr)
	log.Debugf("API address: %v", *apiAddrStr)
	log.Debugf("Threads API address: %v", *threadsAPIAddrStr)

	n, err := threadsCommon.DefaultNetwork(
		*threadsRepo,
		threadsCommon.WithNetHostAddr(hostAddr),
		threadsCommon.WithConnectionManager(connmgr.NewConnManager(*connectionsLowWater, *connectionsHighWater, *connectionsGracePeriod)),
		threadsCommon.WithNetDebug(*debug))
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()
}
