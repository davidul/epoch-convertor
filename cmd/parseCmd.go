package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Long:  "parse",
	Short: "parse",
	Run: func(cmd *cobra.Command, args []string) {
		for i := range args {
			fmt.Println(i)
		}

		getInt64, err := cmd.Flags().GetInt64("millis")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s \n", time.UnixMilli(getInt64).Format(time.RFC850))
		time.Unix(100, 100)
		time.UnixMicro(1000)
	},
}

func init() {
	parseCmd.Flags().Int64("millis", 0, "milliseconds")
	parseCmd.MarkFlagRequired("millis")
	RootCmd.AddCommand(parseCmd)
}
