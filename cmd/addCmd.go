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

		now := t.Now()
		date := now.AddDate(year, month, day)

		//_, err = fmt.Fprintf(cmd.OutOrStdout(), "\033 \u2554 \n")
		//if err != nil {
		//	return
		//}
		//_, err = fmt.Fprintf(cmd.OutOrStdout(), "\u2551 \033[38;5;156m Now %s \n", now.Format(time.UnixDate))
		//if err != nil {
		//	return
		//}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Now %s\n", now.Format(time.UnixDate))
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Add %d year(s) %d month(s) %d day(s)\n", year, month, day)
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", date.Format(time.UnixDate))
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "\t Unix epoch seconds: %d\n", date.Unix())
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "\t Unix epoch miliseconds: %d\n", now.UnixMilli())
		if err != nil {
			return
		}
	},
}

func init() {
	addCmd.Flags().Int("year", 0, "+-year")
	addCmd.Flags().Int("month", 0, "+-month")
	addCmd.Flags().Int("day", 0, "+-day")
	RootCmd.AddCommand(addCmd)
}
