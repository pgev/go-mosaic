package config

import (
	"errors"
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
		IpfsLite:                 defaultIPFSLiteStorePath,
		LogStore:                 defaultLogStorePath,
		ViewStore:                defaultViewStorePath,
		HostAddressString:        "/ip4/0.0.0.0/tcp/4006",
		ConnectionsLowWaterMark:  100,
		ConnectionsHighWaterMark: 400,
		ConnectionsGracePeriod:   20*time.Second,
	}
}

// IpfsLiteStorePath provides the absolute path to the ipfs nodes datastore
// Defaults to "<WorkDir>/threads/ipfslite/"
func (config ThreadsConfig) IpfsLiteStorePath() string {
	return makeAbsolutePath(config.IpfsLite, config.WorkDir)
}

// LogStorePath provides the absolute path to the log store db
// Defaults to "<WorkDir>/threads/logstore/"
func (config ThreadsConfig) LogStorePath() string {
	return makeAbsolutePath(config.LogStore, config.WorkDir)
}

// ViewStorePath provides the absolute path to the view store db
// Defaults to "<WorkDir>/threads/viewstore"
func (config ThreadsConfig) ViewStorePath() string {
	return makeAbsolutePath(config.ViewStore, config.WorkDir)
}

func (config ThreadsConfig) HostAddress() (ma.Multiaddr, error) {
	address, err := ma.NewMultiaddr(config.HostAddressString)
	if err != nil {
		return nil, err
	}
	return address, nil
}

//-----------------------------------------------------------------------------
// Private functions

func (config *ThreadsConfig) validateBasic() error {
	if config.IpfsLite == "" {
		return errors.New("ipfslite_dir cannot be empty")
	}
	if config.LogStore == "" {
		return errors.New("logstore_dir cannot be empty")
	}
	if config.ViewStore == "" {
		return errors.New("viewstore_dir cannot be empty")
	}
	if config.ConnectionsLowWaterMark < 1 {
		return errors.New("connections_low_mark must be strictly positive")
	}
	if config.ConnectionsHighWaterMark < config.ConnectionsLowWaterMark {
		return errors.New("connections_high_mark cannot be smaller than connections_low_mark")
	}
	if config.ConnectionsGracePeriod < 0 {
		return errors.New("connections_grace_period cannot be negative")
	}
	return nil
}
