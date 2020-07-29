package commands

import (
	"fmt"

	logging "github.com/ipfs/go-log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/mosaicdao/go-mosaic/config"
)

var (
	// config is first used to add flags to the commands;
	// on Execute() it is reinitialised to function
	// the config for params to be unmarshalled into.
	config = cfg.DefaultConfig()
	log    = logging.Logger("mosaic")
)

var RootCmd = &cobra.Command{
	Use:   "mosaic",
	Short: "Mosaic node (data stream processing) in Golang",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = parseConfig()
		if err != nil {
			return err
		}

		return nil
	},
}

//-----------------------------------------------------------------------------
// Private functions

func parseConfig() (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	conf.SetWorkDir(conf.WorkDir)
	// cfg.EnsureBasePath(config.BasePath)
	if err = conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %w", err)
	}
	return conf, nil
}
