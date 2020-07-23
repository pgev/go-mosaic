package threads

import (
	cm "github.com/libp2p/go-libp2p-core/connmgr"
)

type NewThreadsOption func(t *threads) error

func WithConnectionManager(connectionManager cm.ConnManager) NewThreadsOption {
	return func(t *threads) error {
		t.connectionManager = connectionManager
		return nil
	}
}
