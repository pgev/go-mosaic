package threads

import (
	logging "github.com/ipfs/go-log"
	"github.com/textileio/go-threads/core/app"
	cm "github.com/libp2p/go-libp2p-core/connmgr"

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

	config            *cfg.ThreadsConfig
	connectionManager cm.ConnManager
}

// NewThreadsNetwork provides a ThreadsNetwork interface to a new instance.
// The ThreadsNetwork must be started by calling Start() before use.
func NewThreadsNetwork(config *cfg.ThreadsConfig) ThreadsNetwork {
	// TODO: take config for OnStart to work

	return &threads{}
}

func (t *threads) OnStart() error {
	// TODO: start ipfslite, logstore etc
	return nil
}

func (t *threads) OnStop() {

}
