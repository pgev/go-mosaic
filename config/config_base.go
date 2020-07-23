package config

import (
	"errors"
)

//-----------------------------------------------------------------------------
// BaseConfig

// BaseConfig defines the base configuration for a Mosaic node
type BaseConfig struct {
	// Working directory for mosaic containing data and config
	WorkDir string `mapstructure:"work_dir"`

	// ConfigDir is the directory name for config path relative to working dir,
	// or it can be given as absolute path.
	// Defaults to ".mosaic/config/"
	ConfigDir string `mapstructure:"config_dir"`

	// ThreadsDir is the directory name for the threads path relative to
	// working dir, or it can be given as absolute path.
	// Defaults to ".mosaic/threads/"
	ThreadsDir string `mapstructure:"threads_dir"`

	// NodePrivateKey is path to JSON file containing node private key
	// relative to working directory, or can be given as absolute path.
	NodePrivateKey string `mapstructure:"node_private_key_file"`
}

// DefaultBaseConfig returns a base configuration with default values.DefaultBaseConfig
// Note WorkDir is left nil, and should be set once with SetWorkDir() before use.
func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		// WorkDir is set in SetWorkDir()
		ConfigDir:      defaultConfigDir,
		ThreadsDir:     defaultThreadsDir,
		NodePrivateKey: defaultNodePrivateKeyPath,
	}
}

// ConfigPath returns the absolute path to the config directory
// Defaults to "<WorkDir>/config"
func (config BaseConfig) ConfigPath() string {
	return makeAbsolutePath(config.ConfigDir, config.WorkDir)
}

// ThreadsPath returns the absolute path to the threads directory
// Defaults to "<WorkDir>/threads"
func (config BaseConfig) ThreadsPath() string {
	return makeAbsolutePath(config.ThreadsDir, config.WorkDir)
}

// NodePrivateKeyFile returns the absolute path to the node's private key file
// Defaults to "<WorkDir>/config/node_private_key.json"
func (config BaseConfig) NodePrivateKeyFile() string {
	return makeAbsolutePath(config.NodePrivateKey, config.WorkDir)
}

//-----------------------------------------------------------------------------
// Private functions

func (config *BaseConfig) validateBasic() error {
	if config.ConfigDir == "" {
		return errors.New("config_dir cannot be empty")
	}
	if config.ThreadsDir == "" {
		return errors.New("threads_dir cannot be empty")
	}
	if config.NodePrivateKey == "" {
		return errors.New("node_private_key_file cannot be empty")
	}
	return nil
}
