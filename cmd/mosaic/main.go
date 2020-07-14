package main

// Special thanks to Tendermint authors for initial inspiration
// on structuring the cli

import (
	cmd "github.com/mosaicdao/go-mosaic/cmd/mosaic/commands"
	"github.com/mosaicdao/go-mosaic/node"
)

func main() {
	rootCmd := cmd.RootCmd

	nodeProvider := node.DefaultNewNode

	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeProvider))

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
