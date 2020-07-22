package threads

//-----------------------------------------------------------------------------
// ThreadsConfig

// ThreadsConfig defines the configuration options for the threads management
type ThreadsConfig struct {
	DataDir                  string        `mapstring:"datadir"`
	NodeKeyFile              string        `mapstructure:"node_key_file"`
	ipfsliteDir              string        `mapstructure:"ipfslite_dir"`
	logstoreDir              string        `mapstructure:"logstore_dir"`
	viewstoreSir             string        `mapstructure:"viewstore_dir`
	HostAddressString        string        `mapstructure:"host_address"`
	ConnectionsLowWaterMark  int           `mapstructure:"connections_low_water_mark"`
	ConnectionsHighWaterMark int           `mapstructure:"connections_high_water_mark"`
	ConnectionsGracePeriod   time.Duration `mapstructure:"connections_grace_period"`
}

func DefaultThreadsConfig() *ThreadsConfig {
	return &ThreadsConfig{}
}
