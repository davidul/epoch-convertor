package cmd

import (
	"epc/models"
	"epc/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"time"
)

var iso map[string]*models.ZoneInfo

var tmzCmd = &cobra.Command{
	Use:   "tz",
	Long:  "tz",
	Short: "tz",
	Run: func(cmd *cobra.Command, args []string) {

		listFlag, err := cmd.Flags().GetBool("list")
		if err != nil {
			fmt.Println(err)
		}

		if listFlag {
			array := pkg.ZoneList(iso)
			sort.Strings(array)
			for _, k := range array {
				fmt.Println(k)
			}
		}

		zoneFromCmd, err2 := cmd.Flags().GetString("timezone")
		if err2 != nil {
			fmt.Println(err2)
		}

		info := iso[zoneFromCmd]

		now := time.Now()
		location, err := time.LoadLocation(info.TimeZone)
		fmt.Fprintf(cmd.OutOrStdout(), location.String())
		if err != nil {
			fmt.Println(err)
		}
		pkg.PrintDate(now.In(location))
	},
}

func init() {
	iso = pkg.ReadZoneInfoISO()
	tmzCmd.Flags().String("timezone", "UTC", "timezone IANA string")
	tmzCmd.Flags().Bool("list", false, "list timezones")
	RootCmd.AddCommand(tmzCmd)
}
