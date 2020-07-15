package main

// Special thanks to Tendermint authors for initial inspiration
// on structuring the cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmd "github.com/mosaicdao/go-mosaic/cmd/mosaic/commands"
	"github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/node"
)

var (
	RootFlag = "root"
)

func main() {
	rootCmd := cmd.RootCmd

	nodeProvider := node.DefaultNewNode

	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeProvider))

	setupRootCommand(rootCmd, "MOSAIC",
		os.ExpandEnv(filepath.Join("$HOME", config.DefaultMosaicDir)))
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic(err)
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func setupRootCommand(cmd *cobra.Command, envPrefix, defaultRoot string) *cobra.Command {
	cmd.PersistentFlags().StringP(RootFlag, "", defaultRoot, "root directory for mosaic data and config")

	return cmd
}

func loadViper() {

}
