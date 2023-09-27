package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var parseCmd = &cobra.Command{
	Use:   "parse unix timestamp with different precisions",
	Short: "parse with different precisions",
	Long: "parse unix timestamp with different precisions \n" +
		"use milliseconds, microseconds or seconds to parse timestamp",
	Run: func(cmd *cobra.Command, args []string) {

		value, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		millis, err := cmd.Flags().GetBool("millis")
		if err != nil {
			fmt.Println(err)
		}

		micro, err := cmd.Flags().GetBool("micro")
		if err != nil {
			fmt.Println(err)
		}

		seconds, err := cmd.Flags().GetBool("seconds")
		if err != nil {
			fmt.Println(err)
		}

		// if none is provided, try all
		if !millis && !micro && !seconds {
			millis = true
			micro = true
			seconds = true
		}

		if millis {
			fmt.Fprintf(cmd.OutOrStdout(), "From millis %s \n", time.UnixMilli(value).Format(time.RFC850))
		}

		if micro {
			fmt.Fprintf(cmd.OutOrStdout(), "From micros %s \n", time.UnixMicro(value).Format(time.RFC850))
		}

		if seconds {
			fmt.Fprintf(cmd.OutOrStdout(), "From seconds %s \n", time.Unix(value, 0).Format(time.RFC850))
		}

	},
}

func init() {
	parseCmd.Flags().Bool("seconds", false, "seconds")
	parseCmd.Flags().Bool("millis", false, "milliseconds")
	parseCmd.Flags().Bool("micro", false, "microseconds")
	RootCmd.AddCommand(parseCmd)
}
