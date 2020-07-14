package main

import (
	"github.com/mosaicdao/go-mosaic/node"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mosaic",
	Short: "Mosaic node (data stream processing) in Golang",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
}

func NewRunNodeCmd(nodeProvider node.NodeProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Run a mosaic node",
		RunE: func(cmd *cobra.Command, args []string) error {
// CONT
		}
	}
}
