package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/* Commands (expected hierarchy)
 *
 * amocli |- tx |- transfer --from <address> --to <address> --amount <number>
 *		  		|- purchase --from <address> --file <hash>
 */

var txCmd = &cobra.Command{
	Use:     "tx",
	Aliases: []string{"t"},
	Short:   "performs a transaction",
	Long:    "amocli tx performs a transaction",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var txTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfers the specified amount of money from <address> to <address>",
	Args:  cobra.NoArgs,
	RunE:  txTransferFunc,
}

var txPurchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "purchases the file specified with file's <hash>",
	Args:  cobra.NoArgs,
	RunE:  txPurchaseFunc,
}

func init() {
	transferCmd := txTransferCmd
	transferCmd.Flags().String("from", "", "specify 'from' address")
	transferCmd.Flags().String("to", "", "specify 'to' address")
	transferCmd.Flags().Uint32("amount", 0, "specify 'amount'")
	transferCmd.MarkFlagRequired("from")
	transferCmd.MarkFlagRequired("to")
	transferCmd.MarkFlagRequired("amount")

	purchaseCmd := txPurchaseCmd
	purchaseCmd.Flags().String("from", "", "specify 'from' address")
	purchaseCmd.Flags().String("file_hash", "", "specify file's hash")
	purchaseCmd.MarkFlagRequired("from")
	purchaseCmd.MarkFlagRequired("file_hash")

	cmd := txCmd
	cmd.AddCommand(transferCmd, purchaseCmd)
}

func txTransferFunc(cmd *cobra.Command, args []string) error {
	var from, to string
	var amount uint32
	var err error

	flags := cmd.Flags()

	if from, err = flags.GetString("from"); err != nil {
		return err
	}
	if to, err = flags.GetString("to"); err != nil {
		return err
	}
	if amount, err = flags.GetUint32("amount"); err != nil {
		return err
	}

	fmt.Printf("'%s' transfers '%d' to '%s'\n", from, amount, to)

	return nil
}

func txPurchaseFunc(cmd *cobra.Command, args []string) error {
	var from, fileHash string
	var err error

	flags := cmd.Flags()

	if from, err = flags.GetString("from"); err != nil {
		return err
	}
	if fileHash, err = flags.GetString("file_hash"); err != nil {
		return err
	}

	fmt.Printf("'%s' purchases '%s'\n", from, fileHash)

	return nil
}
