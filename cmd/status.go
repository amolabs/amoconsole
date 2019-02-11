package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/amolabs/amoconsole/util"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows status of AMO node",
	Long:  "Shows including node info, pubkey, latest block hash, app hash, block height and time",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := util.RPCStatus()
		if err != nil {
			return err
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		fmt.Println(string(resultJSON))

		return nil
	},
}

func init() {
	// init here if needed
}
