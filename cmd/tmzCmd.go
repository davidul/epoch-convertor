package cmd

import (
	"epc/pkg"
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
	"github.com/spf13/cobra"
)

var iso map[string]linkedlist.LinkedList[pkg.ZoneInfo]

var tmzCmd = &cobra.Command{
	Use:   "tz",
	Long:  "tz",
	Short: "tz",
	Run: func(cmd *cobra.Command, args []string) {
		for i := range args {
			fmt.Println(i)
		}

		listFlag, err3 := cmd.Flags().GetBool("list")
		if err3 != nil {
			fmt.Println(err3)
		}

		if listFlag {

			array := pkg.ZoneListSortedArray(iso)
			for _, k := range array {
				fmt.Println(k)
			}
		}

		zoneFromCmd, err2 := cmd.Flags().GetString("timezone")
		if err2 != nil {
			fmt.Println(err2)
		}

		info := iso[zoneFromCmd]
		fmt.Println(info)

		//now := time.Now()
		//location, err := time.LoadLocation(info)
		//fmt.Println(location.String())
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(now.In(location))
	},
}

func init() {
	iso = pkg.ReadZoneInfoISO()
	tmzCmd.Flags().String("timezone", "UTC", "timezone IANA string")
	tmzCmd.Flags().Bool("list", true, "list timezones")
	rootCmd.AddCommand(tmzCmd)
}
