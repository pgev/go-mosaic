package commands

import (
	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config = cfg.DefaultConfig()
)

func ParseConfig() (*cfg.Config, error) {
	config := cfg.DefaultConfig()
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	config
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
