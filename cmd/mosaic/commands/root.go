package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/mosaicdao/go-mosaic/config"
)

var (
	config = cfg.DefaultConfig()
)

func ParseConfig() (*cfg.Config, error) {
	config := cfg.DefaultConfig()
	err := viper.Unmarshal(config) // sets root dir? "--home"
	if err != nil {
		return nil, err
	}
	// config.SetRoot(config.RootDir)
	// cfg.EnsureRoot(config.RootDir)
	if err = config.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "mosaic",
	Short: "Mosaic node (data stream processing) in Golang",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = ParseConfig()
		if err != nil {
			return err
		}

		return nil
	},
}
