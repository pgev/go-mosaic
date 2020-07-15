package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mosaicdao/go-mosaic/node"
)

func NewRunNodeCmd(nodeProvider node.NodeProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Run a mosaic node",
		RunE: func(cmd *cobra.Command, args []string) error {
			// check config
			_, err := nodeProvider(config)
			if err != nil {
				return fmt.Errorf("failed to create new node: %w", err)
			}

			// if err := node.Start(); err != nil {
			// 	return fmt.Errorf("failed to start node: %w", err)
			// }

			// TODO: add node info to log
			log.Infof("Started node")

			// TODO: await
			return nil
		},
	}
	return cmd
}
