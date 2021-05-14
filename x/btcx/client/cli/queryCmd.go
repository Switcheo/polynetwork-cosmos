package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// CmdQueryDenomInfo returns the command that queries a denom info
func CmdQueryDenomInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom-info [denom]",
		Short: "Query denom info",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a specific denom or coin info incluing the coin creator,  coin total supply, the
redeem script and redeem script hash

Example:
$ %s query %s denom-info btcx
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDenomInfoRequest{
				Denom: args[0],
			}

			res, err := queryClient.DenomInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdQueryDenomInfoWithChainId returns the command that queries a denom with chain id
func CmdQueryDenomInfoWithChainId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom-cc-info [denom] [chain_id]",
		Short: "Query denom info correlated with a specific chainID",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a specific denom or coin info correlated with a specific chainID incluing the coin creator,  coin total supply, the
redeem script and redeem script hash, toChainId and the corresponding toAssetHash in hex format

Example:
$ %s query %s denom-cc-info btcx 2
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetDenomCrossChainInfoRequest{
				Denom:   args[0],
				ChainId: chainID,
			}

			res, err := queryClient.DenomCrossChainInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
