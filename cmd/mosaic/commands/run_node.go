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
			n, err := nodeProvider(config)
			if err != nil {
				return fmt.Errorf("failed to create new node: %w", err)
			}

			if err := n.Start(); err != nil {
				return fmt.Errorf("failed to start node: %w", err)
			}

			fmt.Printf("Hello world from %s", n)
			// TODO: add node info to log
			log.Infof("Started node")

			n.Wait()
			return nil
		},
	}
	return cmd
}
