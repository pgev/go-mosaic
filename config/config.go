package config

import (
	"fmt"
	"path/filepath"

	logging "github.com/ipfs/go-log"
)

var (
	// DefaultMosaicDir is used to set the working directory
	// WorkDir will default to "$HOME/.mosaic/"
	DefaultMosaicDir  = ".mosaic"
	defaultConfigDir  = "config"
	defaultThreadsDir = "threads"

	// by default stored under ConfigDir
	defaultConfigFileName     = "config.toml"
	defaultNodePrivateKeyName = "node_private_key.json"

	// DB directories, by default under ThreadsDir
	defaultIpfsLiteStoreDir = "ipfslite"
	defaultLogStoreDir      = "logstore"
	defaultViewStoreDir     = "viewstore"

	defaultConfigFilePath     = filepath.Join(defaultConfigDir,
		defaultConfigFileName)
	defaultNodePrivateKeyPath = filepath.Join(defaultConfigDir,
		defaultNodePrivateKeyName)
	defaultIPFSLiteStorePath  = filepath.Join(defaultThreadsDir, defaultIpfsLiteStoreDir)
	defaultLogStorePath       = filepath.Join(defaultThreadsDir, defaultLogStoreDir)
	defaultViewStorePath      = filepath.Join(defaultThreadsDir, defaultViewStoreDir)

	log = logging.Logger("config")
)

// Config defines the complete configuration for a Mosaic node
type Config struct {
	// Base configuration is at the top-level, unnamed
	BaseConfig `mapstructure:",squash"`

	Threads *ThreadsConfig `mapstructure:"threads"`
}

// DefaultConfig returns a complete configuration with default values set
func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
		Threads:    DefaultThreadsConfig(),
	}
}

// ValidateBasic will return an error for any of the configuration parameters
// which don't pass a basic validation check
func (config *Config) ValidateBasic() error {
	if err := config.BaseConfig.validateBasic(); err != nil {
		return err
	}
	if err := config.Threads.validateBasic(); err != nil {
		return fmt.Errorf("Error in [threads] section: %w", err)
	}
	return nil
}

// SetWorkDir must be called with an absolute path the working directir
func (config *Config) SetWorkDir(workDir string) {
	config.BaseConfig.WorkDir = workDir
	config.Threads.WorkDir = workDir
}

//-----------------------------------------------------------------------------
// Private functions

func makeAbsolutePath(path, abs string) string {
	if abs == "" {
		log.Panic("Absolute path is not set")
	}
	if !filepath.IsAbs(abs) {
		log.Panicf("Provided absolute reference path is not absolute: %s", abs)
	}

	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(abs, path)
}
