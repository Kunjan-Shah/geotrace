package cli

import (
	"strconv"

	"geotrace/x/geotrace/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddLocationProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-location-proof [lat] [lon] [timestamp]",
		Short: "Broadcast message add-location-proof",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLat := args[0]
			argLon := args[1]
			argTimestamp := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddLocationProof(
				clientCtx.GetFromAddress().String(),
				argLat,
				argLon,
				argTimestamp,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
