package threads

import (
	logging "github.com/ipfs/go-log"
	"github.com/textileio/go-threads/core/app"

	"github.com/mosaicdao/go-mosaic/libs/service"
)

var (
	log = logging.Logger("threads")
)

type Threads interface {
}

// Threads provides a Threads network, and implements Servicable and Service.
type threads struct {
	service.BaseService
	app.Net // provides Connection

	threadsDir string
}

func NewThreads() Threads {
	// TODO: take config for OnStart to work
	return &threads{}
}

func OnStart() error {
	// TODO: start ipfslite, logstore etc
	return nil
}

func OnStop() {

}
