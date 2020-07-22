package config

var (
	DefaultMosaicDir  = ".mosaic"
	defaultConfigDir  = "config"
	defaultThreadsDir = "threads"

	defaultConfigFilename = "config.toml"
)

type Config struct {
	// Base configuration is at the top-level, unnamed
	BaseConfig `mapstructure:",squash"`

	Column  *ColumnConfig  `mapstructure:"column"`
	Threads *ThreadsConfig `mapstructure:"threads"`
}

func DefaultConfig() *Config {
	return &Config{
		BaseConfig: DefaultBaseConfig(),
		Column:     DefaultColumnConfig(),
	}
}

func (config *Config) ValidateBasic() error {
	// TODO: all validation for subsections
	return nil
}
