package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	p2ppeerstore "github.com/libp2p/go-libp2p-core/peerstore"

	"github.com/mosaicdao/go-mosaic/node"
)

func AddNodeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("threads.host", "a", config.Threads.HostAddressString, "Host address")
}

func NewRunNodeCmd(nodeProvider node.NodeProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Run a mosaic node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			node, landscape, err := nodeProvider(ctx, config)
			if err != nil {
				return fmt.Errorf("failed to create new node: %w", err)
			}

			if err := node.Start(); err != nil {
				return fmt.Errorf("failed to start node: %w", err)
			}
			log.Info("Started node")

			// TODO: catch SIGTERM interrupt
			go func() {
				for {
					time.Sleep(5 * time.Second)
					fmt.Printf("connected peers (%v)\n",
						node.Threads().Host().Network().Peers(),
					)
				}
			}()

			go func() {
				for {
					time.Sleep(2 * time.Second)
					members := landscape.Peers(nil)
					for _, member := range members {
						if member.ID != node.Threads().Host().ID() {
							fmt.Printf("Adding peer %s\n", member.ID.String())
							node.Threads().Host().Peerstore().AddAddrs(
								member.ID,
								member.Addrs,
								p2ppeerstore.PermanentAddrTTL,
							)
						}
					}
				}
			}()
			node.Wait()
			return nil
		},
	}

	AddNodeFlags(cmd)
	return cmd
}
