package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var parseCmd = &cobra.Command{
	Use:   "parse unix timestamp with different precisions",
	Short: "parse with different precisions",
	Long: "parse unix timestamp with different precisions \n" +
		"use milliseconds, microseconds or seconds to parse timestamp",
	Run: func(cmd *cobra.Command, args []string) {

		millis, err := cmd.Flags().GetInt64("millis")
		if err != nil {
			fmt.Println(err)
		}
		if millis != -1 {
			fmt.Fprintf(cmd.OutOrStdout(), "%s \n", time.UnixMilli(millis).Format(time.RFC850))
		}

		micro, err := cmd.Flags().GetInt64("micro")
		if err != nil {
			fmt.Println(err)
		}

		if micro != -1 {
			fmt.Fprintf(cmd.OutOrStdout(), "%s \n", time.UnixMicro(micro).Format(time.RFC850))
		}

		seconds, err := cmd.Flags().GetInt64("seconds")
		if err != nil {
			fmt.Println(err)
		}
		if seconds != -1 {
			fmt.Fprintf(cmd.OutOrStdout(), "%s \n", time.Unix(seconds, 0).Format(time.RFC850))
		}

	},
}

func init() {
	parseCmd.Flags().Int64("seconds", -1, "seconds")
	parseCmd.Flags().Int64("millis", -1, "milliseconds")
	parseCmd.Flags().Int64("micro", -1, "microseconds")
	RootCmd.AddCommand(parseCmd)
}
