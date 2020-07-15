package main

// Special thanks to Tendermint authors for initial inspiration
// on structuring the cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmd "github.com/mosaicdao/go-mosaic/cmd/mosaic/commands"
	"github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/node"
)

var (
	BasePathFlag = "base_path"
)

func main() {
	rootCmd := cmd.RootCmd

	nodeProvider := node.DefaultNewNode

	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeProvider))

	setupRootCommand(rootCmd, "MOSAIC",
		os.ExpandEnv(filepath.Join("$HOME", config.DefaultMosaicDir)))

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func setupRootCommand(cmd *cobra.Command, envPrefix, defaultBasePath string) *cobra.Command {
	cmd.PersistentFlags().StringP(BasePathFlag, "d", defaultBasePath, "base path for mosaic data and config")

	initEnvironment(envPrefix) // TODO: ? wrap in cobra.OnInitialize
	// bind all flags to viper
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
	cmd.PersistentPreRunE = concatCobraCmdFuncs(loadViper, cmd.PersistentPreRunE)

	return cmd
}

func initEnvironment(envPrefix string) {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
}

func loadViper(cmd *cobra.Command, args []string) error {
	basePath := viper.GetString(BasePathFlag)
	// fix base path to decouple from other configuration sources
	viper.Set(BasePathFlag, basePath)
	viper.SetConfigName("config")
	// search in base path and /config subdirectory
	viper.AddConfigPath(basePath)
	viper.AddConfigPath(filepath.Join(basePath, "config"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// ignore ConfigFileNotFound error
		} else {
			// return error if config file is found but another error occured
			return err
		}
	}
	return nil
}

// From github.com/tendermint/tendermint
type cobraCmdFunc func(cmd *cobra.Command, args []string) error

// Returns a single function that calls each argument function in sequence
// RunE, PreRunE, PersistentPreRunE, etc. all have this same signature
func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range fs {
			if f != nil {
				if err := f(cmd, args); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
