package column

import (
	"github.com/mosaicdao/go-mosaic/types"
)

type Member interface {
	getAddress() types.Address
}

type Member struct {
}
