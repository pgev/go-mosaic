package config

import (
	"time"

	ma "github.com/multiformats/go-multiaddr"
)

//-----------------------------------------------------------------------------
// ThreadsConfig

// ThreadsConfig defines the configuration options for the threads management
type ThreadsConfig struct {
	WorkDir                  string        `mapstring:"work_dir"`

	// Path to IPFSLite datastore, relative to working dir or absolute
	IpfsLite                 string        `mapstructure:"ipfslite_dir"`
	// Path to Log datastore, relative to working dir or absolute
	LogStore                 string        `mapstructure:"logstore_dir"`
	// Path to View datastore, relative to working dir or absolute
	ViewStore                string        `mapstructure:"viewstore_dir`

	HostAddressString        string        `mapstructure:"host_address"`
	ConnectionsLowWaterMark  int           `mapstructure:"connections_low_mark"`
	ConnectionsHighWaterMark int           `mapstructure:"connections_high_mark"`
	ConnectionsGracePeriod   time.Duration `mapstructure:"connections_grace_period"`
}

// DefaultThreadsConfig returns a configuration for Threads with default values.
// Note WorkDir is left nil, and should be set once with SetWorkDir() before use.
func DefaultThreadsConfig() *ThreadsConfig {
	return &ThreadsConfig{
		// WorkDir is set in SetWorkDir()
		IpfsLite: defaultIPFSLiteStorePath,

	}
}

// IpfsLiteStorePath provides the absolute path to the ipfs nodes datastore
// Defaults to ".mosaic/threads/ipfslite/"
func (cfg ThreadsConfig) IpfsLiteStorePath() string {
	return makeAbsolutePath(cfg.IpfsLite, cfg.WorkDir)
}

// LogStorePath provides the absolute path to the log store db
// Defaults to ".mosaic/threads/logstore/"
func (cfg ThreadsConfig) LogStorePath() string {
	return makeAbsolutePath(cfg.LogStore, cfg.WorkDir)
}

// ViewStorePath provides the absolute path to the view store db
// Defaults to ".mosaic/threads/viewstore"
func (cfg ThreadsConfig) ViewStorePath() string {
	return makeAbsolutePath(cfg.ViewStore, cfg.WorkDir)
}

func (cfg ThreadsConfig) HostAddress() (ma.Multiaddr, error) {
	address, err := ma.NewMultiaddr(cfg.HostAddressString)
	if err != nil {
		return nil, err
	}
	return address, nil
}
