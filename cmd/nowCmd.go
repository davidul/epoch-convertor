package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "now",
	Long:  "now displays current time in unix epoch",
	RunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now()

		_, err := fmt.Fprintf(cmd.OutOrStdout(), "Current time is: %s \n", now.Format("2006-01-02T15:04:05"))
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch seconds: %d \n", now.Unix())
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch miliseconds: %d \n", now.UnixMilli())
		if err != nil {
			return err
		}

		return err
	},
}

func init() {
	RootCmd.AddCommand(nowCmd)
}
