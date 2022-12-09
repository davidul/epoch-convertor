package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add year/month/date. Can be positive or negative value.",
	Short: "add year/month/date",
	Long:  "add long",
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Printf("\033 \u2554 \n")
		fmt.Printf("\u2551 \033[38;5;156m Now %s \n", now.Format(time.UnixDate))
		fmt.Printf("Add %d year(s) %d month(s) %d day(s) \n", year, month, day)
		fmt.Printf("%s \n", date.Format(time.UnixDate))
		fmt.Printf("\t Unix epoch seconds: %d \n", date.Unix())
		fmt.Printf("\t Unix epoch miliseconds: %d \n", now.UnixMilli())
	},
}

func init() {
	addCmd.Flags().Int("year", 0, "+-year")
	addCmd.Flags().Int("month", 0, "+-month")
	addCmd.Flags().Int("day", 0, "+-day")
}
