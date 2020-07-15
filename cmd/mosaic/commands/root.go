package commands

import (
	logging "github.com/ipfs/go-log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/mosaicdao/go-mosaic/config"
)

var (
	config = cfg.DefaultConfig()
	log    = logging.Logger("mosaic")
)

func ParseConfig() (*cfg.Config, error) {
	config := cfg.DefaultConfig()
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	// config.SetBasePath(config.BasePath)
	// cfg.EnsureBasePath(config.BasePath)
	// if err = config.ValidateBasic(); err != nil {
	// 	return nil, fmt.Errorf("error in config file: %v", err)
	// }
	return config, nil
}

var RootCmd = &cobra.Command{
	Use:   "mosaic",
	Short: "Mosaic node (data stream processing) in Golang",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = ParseConfig()
		if err != nil {
			return err
		}
		//TODO : set up logging
		return nil
	},
}
