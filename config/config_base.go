package config

import (
	"errors"

	moos "github.com/mosaicdao/go-mosaic/libs/os"
)

//-----------------------------------------------------------------------------
// BaseConfig

// BaseConfig defines the base configuration for a Mosaic node
type BaseConfig struct {
	// Working directory for mosaic containing data and config
	WorkDir string `mapstructure:"work_dir"`

	// NodePrivateKey is path to a file containing node private key
	// relative to working directory, or can be given as absolute path.
	NodePrivateKey string `mapstructure:"node_private_key_file"`
}

// DefaultBaseConfig returns a base configuration with default values.DefaultBaseConfig
// Note WorkDir is left nil, and should be set once with SetWorkDir() before use.
func defaultBaseConfig() BaseConfig {
	return BaseConfig{
		// WorkDir is set in SetWorkDir()
		NodePrivateKey: defaultNodePrivateKeyPath,
	}
}

// ConfigPath provides the absolute path to the config directory
// Defaults to "<WorkDir>/config/"
func (config BaseConfig) ConfigPath() string {
	return makeAbsolutePath(DefaultConfigDir, config.WorkDir)
}

// NodePrivateKeyFile returns the absolute path to the node's private key file
// Defaults to "<WorkDir>/config/node_private_key.json"
func (config BaseConfig) NodePrivateKeyFile() string {
	return makeAbsolutePath(config.NodePrivateKey, config.WorkDir)
}

//-----------------------------------------------------------------------------
// Private functions

func (config *BaseConfig) validateBasic() error {
	if config.NodePrivateKey == "" {
		return errors.New("node_private_key_file cannot be empty")
	}
	return nil
}

func (config *BaseConfig) ensurePaths() error {
	// make workdir
	if err := moos.EnsureDir(config.WorkDir); err != nil {
		return err
	}
	// make config dir
	if err := moos.EnsureDir(config.ConfigPath()); err != nil {
		return err
	}
	return nil
}
