package threads

import (
	ma "github.com/multiformats/go-multiaddr"
	cm "github.com/libp2p/go-libp2p-core/connmngr"
	logging "github.com/ipfs/go-log"
	"github.com/textileio/go-threads/core/app"

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
	app.Net // provides Connection

	config *cfg.ThreadsConfig
	threadsDir  string // store Mosaic db, default
	ipfsLiteDir string //
}

// NewThreadsNetwork provides a ThreadsNetwork interface.
// The ThreadsNetwork must be started by calling Start() before use.
// Required options are :
//  - WithHostAddress
//  - WithConnectionManager
func NewThreadsNetwork(
	config *cfg.ThreadsConfig
	opts ..NewThreadsOption
) ThreadsNetwork {

	// TODO: take config for OnStart to work
	config := &NewThreadsConfig{}
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if config.HostAddress == nil {
		return nil, err
	}

	if config.ConnectionManager == nil {
		return nil, err
	}


	return &threads{}
}

func (t *threads) OnStart() error {
	// TODO: start ipfslite, logstore etc
	return nil
}

func (t *threads) OnStop() {

}
