package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/KiraCore/sekai/x/tokens/types"
)

// GetCmdQueryTokenAlias the query delegation command.
func GetCmdQueryTokenAlias() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias symbol",
		Short: "Get the token alias by symbol",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.TokenAliasRequest{Symbol: args[0]}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetTokenAlias(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Data)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}