package cli

import (
	"fmt"
	"os"

	"filippo.io/age"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"code.zeeve.net/client-projects/iopn/v2/x/e2ee/keyring"
	"code.zeeve.net/client-projects/iopn/v2/x/e2ee/types"
)

const FlagKeyringName = "keyring-name"

func KeygenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keygen",
		Short: "Generates a new native X25519 key pair",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			krName, err := cmd.Flags().GetString(FlagKeyringName)
			if err != nil {
				return err
			}

			kr, err := keyring.New("cronosd", clientCtx.Keyring.Backend(), clientCtx.HomeDir, os.Stdin)
			if err != nil {
				return err
			}

			k, err := age.GenerateX25519Identity()
			if err != nil {
				return err
			}

			if err := kr.Set(krName, []byte(k.String())); err != nil {
				return err
			}

			fmt.Println(k.Recipient())
			return nil
		},
	}

	cmd.Flags().String(FlagKeyringName, types.DefaultKeyringName, "The keyring name to use")

	return cmd
}
