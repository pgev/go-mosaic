package config

// BaseConfig defines the base configuration for a Mosaic node
type BaseConfig struct {
	// Root directory for mosaic containing all data
	RootDir string `mapstructure:"root_dir"`

	// Path to JSON file containing Member private key
	MemberPrivateKeyFile string `mapstructure:"member_private_key_file"`

	// CONTINUE
	ColumnThreadId string `mapstructure:"column_thread_id"`
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{}
}
