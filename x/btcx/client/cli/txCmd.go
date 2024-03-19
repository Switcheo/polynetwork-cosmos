package cli

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	sdkmath "cosmossdk.io/math"
	"github.com/spf13/cobra"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
)

func CmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [denom] [redeemScript]",
		Short: "Creates a new denom for the given redeem script",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s tx %s create btc-111 12345678
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenom := string(args[0])
			argsRedeemScript := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreate(clientCtx.GetFromAddress().String(), string(argsDenom), string(argsRedeemScript))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBind() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bind [denom] [to_chain_id] [to_asset_hash]",
		Short: "Binds an asset hash to a source denom. Must be an operator.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s tx %s bind btc-222 3 12341234
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			sourceAssetDenom := args[0]

			toChainId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			toAssetHashStr := args[2]
			if toAssetHashStr[0:2] == "0x" {
				toAssetHashStr = toAssetHashStr[2:]
			}

			toAssetHash, err := hex.DecodeString(toAssetHashStr)
			if err != nil {
				return fmt.Errorf("decode hex string 'toAssetHash' error:%v", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgBind(clientCtx.GetFromAddress().String(), sourceAssetDenom, toChainId, toAssetHash)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock [denom] [to_chain_id] [to_address] [amount]",
		Short: "Locks the amount of denom and releases the same amount on to_chain_id chain to to_address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s tx %s lock btc-xxx 3 616f2a4a38396ff203ea01e6c070ae421bb8ce2d 123
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			sourceAssetDenom := args[0]

			toChainIdStr := args[1]
			toChainId, err := strconv.ParseUint(toChainIdStr, 10, 64)
			if err != nil {
				return err
			}

			toAddressStr := args[2]
			toAddress, err := hex.DecodeString(toAddressStr)
			if err != nil {
				return fmt.Errorf("decode hex string 'toAddress' error:%v", err)
			}

			valueBigInt, ok := big.NewInt(0).SetString(args[3], 10)
			if !ok {
				return fmt.Errorf("read value as big int from args[3] failed")
			}
			value := sdkmath.NewIntFromBigInt(valueBigInt)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLock(clientCtx.GetFromAddress().String(), sourceAssetDenom, toChainId, toAddress, value)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
