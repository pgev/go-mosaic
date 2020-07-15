package config

type ColumnConfig struct {
	ThreadId string `mapstructure:"thread_id"`
}

var ()

func DefaultColumnConfig() *ColumnConfig {
	return &ColumnConfig{}
}

func TestColumnConfig() *ColumnConfig {
	return &ColumnConfig{}
}
