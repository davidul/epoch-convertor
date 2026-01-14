package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var unixCmd = &cobra.Command{
	Use:   "unix  - converts unix timestamp with different precisions",
	Short: "unix converts unix timestamp with different precisions",
	Long: "unix timestamp with different precisions \n" +
		"use milliseconds, microseconds or seconds to parse timestamp",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("Please provide unix timestamp")
			return
		}
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
	unixCmd.Flags().Bool("seconds", false, "seconds")
	unixCmd.Flags().Bool("millis", false, "milliseconds")
	unixCmd.Flags().Bool("micro", false, "microseconds")
	RootCmd.AddCommand(unixCmd)
}
