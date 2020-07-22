package threads

import (
	ma "github.com/multiformats/go-multiaddr"
	cm "github.com/libp2p/go-libp2p-core/connmngr"
)

type NewThreadsConfig struct {
	HostAddress       ma.Multiaddr
	ConnectionManager cm.ConnManager

}

type NewThreadsOption func(config *NewThreadsConfig) error

func WithHostAddress(address *ma.Multiaddr) NewThreadsOption {
	return func(config *NewThreadsConfig) error {
		config.HostAddress = address
		return nil
	}
}

func WithConnectionManager(connectionManager cm.ConnManager) NewThreadsOption {
	return func(config *NewThreadsConfig) error {
		config.ConnectionManager = connectionManager
		return nil
	}
}
