package config

import (
	"path/filepath"

	logging "github.com/ipfs/go-log"
)

var (
	// DefaultMosaicDir is taken as the working directory relative
	// to the provided working directory, eg. "$HOME/.mosaic/"
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

	defaultConfigFilePath     = filepath.Join(defaultConfigDir, defaultConfigFileName)
	defaultNodePrivateKeyPath = filepath.Join(defaultConfigDir, defaultNodePrivateKeyName)
	defaultIPFSLiteStorePath  = filepath.Join(defaultThreadsDir, defaultIpfsLiteStoreDir)
	defaultLogStorePath       = filepath.Join(defaultThreadsDir, defaultLogStoreDir)
	defaultViewStorePath      = filepath.Join(defaultThreadsDir, defaultViewStoreDir)

	log = logging.Logger("config")
)

type Config struct {
	// Base configuration is at the top-level, unnamed
	BaseConfig `mapstructure:",squash"`

	Threads *ThreadsConfig `mapstructure:"threads"`
}

func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
		Threads:    DefaultThreadsConfig(),
	}
}

func (config *Config) ValidateBasic() error {
	// TODO: all validation for subsections
	return nil
}

func (config *Config) SetWorkDir(workDir string) {
	config.BaseConfig.WorkDir = workDir
	config.Threads.WorkDir = workDir
}
//-----------------------------------------------------------------------------
//

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
