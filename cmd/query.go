package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	atypes "github.com/amolabs/amoabci/amo/types"
	"github.com/amolabs/amoconsole/tx"
)

/* Commands (expected hierarchy)
 *
 * amocli |- query
 */

var queryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"q"},
	Short:   "Performs a query ...",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var queryAddressCmd = &cobra.Command{
	Use:   "address [address]",
	Short: "Shows address's general information",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := atypes.NewAddressFromBytes([]byte(args[0]))
		targetInfo, err := tx.QueryAddressInfo(*target)
		if err != nil {
			return err
		}

		fmt.Println(string(targetInfo))

		return nil
	},
}

func init() {
	// init here if needed
	addressCmd := queryAddressCmd
	cmd := queryCmd
	cmd.AddCommand(addressCmd)
}
