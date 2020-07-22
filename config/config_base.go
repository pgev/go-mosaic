package config

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
	// Defaults to ".mosaic/config/node_private_key.json"
	NodePrivateKey string `mapstructure:"member_private_key_file"`
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		// WorkDir is set in SetWorkDir()
		ConfigDir:      defaultConfigDir,
		ThreadsDir:     defaultConfigDir,
		NodePrivateKey: defaultNodePrivateKeyPath,
	}
}

func (config BaseConfig) ConfigPath() string {
	return makeAbsolutePath(config.ConfigDir, config.WorkDir)
}

func (config BaseConfig) ThreadsPath() string {
	return makeAbsolutePath(config.ThreadsDir, config.WorkDir)
}

func (config BaseConfig) NodePrivateKeyFile() string {
	return makeAbsolutePath(config.NodePrivateKey, config.WorkDir)
}
