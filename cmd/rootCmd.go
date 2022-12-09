package cmd

import (
	"epc/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "epc",
	Short: "epc",
	Long:  "epc",
	Run: func(cmd *cobra.Command, args []string) {
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Println(err)
		}

		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Println(err)
		}
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			fmt.Println(err)
		}
		day, err := cmd.Flags().GetInt("day")
		if err != nil {
			fmt.Println(err)
		}

		now := time.Now()
		date := now.AddDate(year, month, day)

		millis, err := cmd.Flags().GetBool("millis")
		if err != nil {
			fmt.Println(err)
		}

		seconds, err := cmd.Flags().GetBool("seconds")
		if err != nil {
			fmt.Println(err)
		}

		if format != "" {
			fmt.Printf("Time in custom format %s\n", pkg.PrintTimeFormatted(date, format))
		}

		pkg.PrintUnix(date, millis, seconds)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.Flags().String("format", time.RFC3339, "format of date time")
	rootCmd.Flags().Int("year", 0, "+/-year")
	rootCmd.Flags().Int("month", 0, "+/-month")
	rootCmd.Flags().Int("day", 0, "+/-day")
	rootCmd.Flags().Bool("millis", false, "unix time in millis")
	rootCmd.Flags().Bool("seconds", true, "unix time in seconds")
}
