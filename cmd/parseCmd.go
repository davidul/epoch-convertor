package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var parseCmd = &cobra.Command{
	Use:   "parse date/time in different formats",
	Long:  "parse",
	Short: "parse.",
	Run: func(cmd *cobra.Command, args []string) {
		date, err := cmd.Flags().GetString("date")
		fmt.Println("Date", date)
		if err != nil {
			fmt.Println(err)
		}

		format, err := cmd.Flags().GetString("in")
		if err != nil {
			fmt.Println(err)
		}
		if format == "" {
			mapString := viper.GetStringMapString("date-formats")
			format = mapString["default-format"]
		}

		parse, err := time.Parse(format, date)
		if err != nil {
			fmt.Println(err)
		}

		//out, err := cmd.Flags().GetString("out")
		//if err != nil {
		//	fmt.Println(err)
		//}

		s := parse.Format(time.Stamp)
		fmt.Println(s)
		fmt.Println(parse.Unix())
	},
}

func init() {
	parseCmd.Flags().String("in", "", "in")
	parseCmd.Flags().String("date", "2006-01-02T15:04:05", "date")
	parseCmd.Flags().String("out", "", "out")
	RootCmd.AddCommand(parseCmd)
}
