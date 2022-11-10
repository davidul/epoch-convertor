package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add",
	Long:  "add",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
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
		fmt.Printf("Now %s \n", now.Format(time.UnixDate))
		fmt.Printf("Now %d %s \n", year, date.Format(time.UnixDate))
		fmt.Printf("Unix epoch seconds: %d \n", date.Unix())
		fmt.Printf("Unix epoch miliseconds: %d \n", now.UnixMilli())
	},
}

func init() {
	addCmd.Flags().Int("year", 0, "")
	addCmd.Flags().Int("month", 0, "")
	addCmd.Flags().Int("day", 0, "")
	rootCmd.AddCommand(addCmd)
}
