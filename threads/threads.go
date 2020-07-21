package threads

import (
	ma "github.com/multiformats/go-multiaddr"
	cm "github.com/libp2p/go-libp2p-core/connmngr"
	logging "github.com/ipfs/go-log"
	"github.com/textileio/go-threads/core/app"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

var (
	log = logging.Logger("threads")
)

type ThreadsNetwork interface {
}

// Threads provides a Threads network, and implements Servicable and Service.
type threads struct {
	service.BaseService
	app.Net // provides Connection

	threadsDir string
}

func NewThreadsNetwork(
	threadsDir string,
	hostAddress *ma.Multiaddr,
	connectionManager cm.ConnManager
	opts ...NewThreadsOption
) ThreadsNetwork {
	// TODO: take config for OnStart to work
	config := &NewThreadsConfig{}
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}



	return &threads{}
}

func (t *threads) OnStart() error {
	// TODO: start ipfslite, logstore etc
	return nil
}

func (t *threads) OnStop() {

}
