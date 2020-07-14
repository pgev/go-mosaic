package main

import (
	"github.com/mosaicdao/go-mosaic/node"
	"github.com/spf13/cobra"
)


func NewRunNodeCmd(nodeProvider node.NodeProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Run a mosaic node",
		RunE: func(cmd *cobra.Command, args []string) error {
			// check config


		}
	}
}
