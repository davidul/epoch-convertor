package cmd

import (
	"epc/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var RootCmd = &cobra.Command{
	Use:   "epc",
	Short: "epc",
	Long:  "epc outputs current time. You can add/subtract day, month and year to it.",
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

var t pkg.StaticTime

func SetStaticTime(time pkg.StaticTime) {
	t = time
}
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	RootCmd.Flags().String("format", time.RFC3339, "format of date time")
	RootCmd.Flags().Int("year", 0, "+/-year")
	RootCmd.Flags().Int("month", 0, "+/-month")
	RootCmd.Flags().Int("day", 0, "+/-day")
	RootCmd.Flags().Bool("millis", false, "unix time in millis")
	RootCmd.Flags().Bool("seconds", true, "unix time in seconds")

}
