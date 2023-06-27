package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var stopwatchCmd = &cobra.Command{
	Use:   "stopwatch",
	Long:  "stopwatch",
	Short: "stopwatch",
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(time.Second)

		seconds := 0
		minutes := 0
		hours := 0

		secStr := ""
		minStr := ""
		hourStr := ""

		for _ = range ticker.C {
			seconds++
			if seconds == 60 {
				seconds = 0
				minutes++
			}

			if minutes == 60 {
				minutes = 0
				hours++
			}
			if seconds < 10 {
				secStr = "0" + strconv.Itoa(seconds)
			} else {
				secStr = strconv.Itoa(seconds)
			}

			if minutes < 10 {
				minStr = "0" + strconv.Itoa(minutes)
			} else {
				minStr = strconv.Itoa(minutes)
			}

			if hours < 10 {
				hourStr = "0" + strconv.Itoa(hours)
			} else {
				hourStr = strconv.Itoa(hours)
			}
			//fmt.Printf("\033[0;0H")
			fmt.Fprintf(cmd.OutOrStdout(), "\033[2K\r")
			fmt.Fprintf(cmd.OutOrStdout(), "%s:%s:%s", hourStr, minStr, secStr)
		}

	},
}

func init() {
	RootCmd.AddCommand(stopwatchCmd)
}
