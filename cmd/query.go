package cmd

import (
	"github.com/spf13/cobra"
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

func init() {
	// init here if needed
}
