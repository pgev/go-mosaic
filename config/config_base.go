package config

// BaseConfig defines the base configuration for a Mosaic node
type BaseConfig struct {
	// Working directory for mosaic containing data and config
	WorkDir string `mapstructure:"work_dir"`

	//
	ConfigDir string `mapstructure:"config_dir"`

	//
	ThreadsDir string `mapstructure:"threads_dir"`

	// Path to JSON file containing Member private key
	MemberPrivateKeyFile string `mapstructure:"member_private_key_file"`
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{}
}
