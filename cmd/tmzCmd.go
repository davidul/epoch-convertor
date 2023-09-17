package cmd

import (
	"epc/models"
	"epc/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strconv"
	"strings"
	"time"
)

var iso map[string]*models.ZoneInfo

var list []*models.TimeZoneInfo

var zoneMap map[string]*models.TimeZoneInfo

var tmzCmd = &cobra.Command{
	Use:   "tz to list and select timezone",
	Long:  "tz long",
	Short: "tz",
	Run: func(cmd *cobra.Command, args []string) {

		filter := cmd.Flag("filter")

		listFlag, err := cmd.Flags().GetBool("list")
		if err != nil {
			fmt.Println(err)
		}

		if listFlag {
			array := pkg.ZoneList(iso)
			sort.Strings(array)
			for _, k := range array {
				if filter != nil && filter.Value.String() != "" {
					strFilter := filter.Value.String()
					if strings.Contains(k, strFilter) {
						fmt.Println(k)
						fmt.Println(iso[k].CountryCode)
					}
				} else {
					fmt.Println(k)
				}
			}
			return
		}

		zoneFromCmd, err2 := cmd.Flags().GetString("timezone")
		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(zoneFromCmd)
		if zoneFromCmd != "" {
			info := zoneMap[zoneFromCmd]
			//now := time.Now()
			location, err := time.LoadLocation(zoneFromCmd)
			fmt.Fprintf(cmd.OutOrStdout(), "Name: %s\n", info.Name)
			fmt.Fprintf(cmd.OutOrStdout(), "Location: %s\n", location.String())
			fmt.Fprintf(cmd.OutOrStdout(), "Country Code: %s \n", info.ZoneInfo.CountryCode)
			fmt.Fprintf(cmd.OutOrStdout(), "Offset: %v\n", strconv.Itoa(info.Offset))
			if err != nil {
				fmt.Println(err)
			}
			//pkg.PrintDate(now.In(location))

			return
		}
	},
}

func init() {
	iso = pkg.ReadZoneInfoISO()
	list = pkg.TimeZoneList()
	zoneMap = pkg.TimeZoneMap()
	tmzCmd.Flags().String("timezone", "Europe/London", "timezone IANA string")
	tmzCmd.Flags().String("filter", "", "filter timezones")
	tmzCmd.Flags().Bool("list", false, "list timezones")
	RootCmd.AddCommand(tmzCmd)
}

func zone(tz string) {

}

func zoneByTime(t int) {

}
