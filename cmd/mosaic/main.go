package main

import (
	// cmd "github.com/mosaicdao/go-mosaic/cmd/mosaic/commands"
	"github.com/mosaicdao/go-mosaic/node"
)

func main() {
	rootCmd := RootCmd

	nodeProvider := node.DefaultNewNode

	rootCmd.AddCommand(NewRunNodeCmd(nodeProvider))

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
