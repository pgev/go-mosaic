package config

var (
	DefaultMosaicDir = ".mosaic"

	defaultConfigFilename = "config.toml"
)

type Config struct {
	// Base configuration is at the top-level, unnamed
	BaseConfig `mapstructure:",squash"`

	Column *ColumnConfig `mapstructure:"column"`
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
