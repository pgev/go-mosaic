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
	cfg "github.com/mosaicdao/go-mosaic/config"
	"github.com/mosaicdao/go-mosaic/node"
)

var (
	workPathFlag = "work_dir"
)

func main() {
	rootCmd := cmd.RootCmd

	// nodeProvider := node.DefaultNewNode
	nodeProvider := node.CutCornersNewNode

	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeProvider))

	c := setupRootCommand(rootCmd, "MOSAIC",
		os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultMosaicDir)))

	if err := c.Execute(); err != nil {
		panic(err)
	}
}

func setupRootCommand(c *cobra.Command, envPrefix, defaultWorkPath string) *cobra.Command {
	cobra.OnInitialize(func() { initEnvironment(envPrefix) })
	c.PersistentFlags().StringP(workPathFlag, "d", defaultWorkPath, "work path for mosaic data and config")
	c.PersistentPreRunE = concatCobraCmdFuncs(setupViper, c.PersistentPreRunE)

	return c
}

func initEnvironment(envPrefix string) {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
}

func setupViper(cmd *cobra.Command, args []string) error {
	// bind all flags to viper
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	workPath := viper.GetString(workPathFlag)
	// fix work path to decouple from other configuration sources
	viper.Set(workPathFlag, workPath)
	viper.SetConfigName("config")
	// search in work path and /config subdirectory
	viper.AddConfigPath(workPath)
	viper.AddConfigPath(filepath.Join(workPath, cfg.DefaultConfigDir))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
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
