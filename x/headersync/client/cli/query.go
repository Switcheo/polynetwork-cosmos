package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	// "strings"

	"github.com/gogo/protobuf/codec"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group headersync queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}

// GetCmdQueryConsensusPeers queries the current consensus peers for the given chain id.
func GetCmdQueryConsensusPeers(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "consensus-peers [chainID]",
		Args:  cobra.ExactArgs(2),
		Short: "Query the consensus peers for a specific chainID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the current consensus peers info for a specific chainID that is
already synced from another blockchain into current chain. Typically this would be the polynetwork chain (with chainID=0).

Example:
$ %s query %s consensus-peers 0
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			chainIDStr := args[0]
			chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetConsensusPeersRequest{
				ChainId: chainID,
			}

			res, err := queryClient.ConsensusPeers(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}
