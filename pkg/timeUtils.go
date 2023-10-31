package pkg

import (
	"epc/models"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func PrintFormats(t time.Time, cmd *cobra.Command) {

	fmt.Fprintf(cmd.OutOrStdout(), "Reference    %s \n", t.Format(time.Layout))
	fmt.Fprintf(cmd.OutOrStdout(), "ANSIC    %s \n", t.Format(time.ANSIC))
	fmt.Fprintf(cmd.OutOrStdout(), "Kitchen  %s \n", t.Format(time.Kitchen))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC822   %s \n", t.Format(time.RFC822))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC822Z %s \n", t.Format(time.RFC822Z))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC850   %s \n", t.Format(time.RFC850))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC1123  %s \n", t.Format(time.RFC1123))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC1123Z %s \n", t.Format(time.RFC1123Z))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC3339  %s \n", t.Format(time.RFC3339))
	fmt.Fprintf(cmd.OutOrStdout(), "RFC3339Nano %s \n", t.Format(time.RFC3339Nano))
	fmt.Fprintf(cmd.OutOrStdout(), "UnixDate %s \n", t.Format(time.UnixDate))
	fmt.Fprintf(cmd.OutOrStdout(), "RubyDate %s \n", t.Format(time.RubyDate))
	fmt.Fprintf(cmd.OutOrStdout(), "Stamp %s \n", t.Format(time.Stamp))
	fmt.Fprintf(cmd.OutOrStdout(), "Stamp Milli %s \n", t.Format(time.StampMilli))
	fmt.Fprintf(cmd.OutOrStdout(), "Stamp Micro %s \n", t.Format(time.StampMicro))
	fmt.Fprintf(cmd.OutOrStdout(), "Stamp Nano %s \n", t.Format(time.StampNano))
	fmt.Fprintf(cmd.OutOrStdout(), "DateOnly %s \n", t.Format(time.DateOnly))
	fmt.Fprintf(cmd.OutOrStdout(), "DateOnly %s \n", t.Format(time.DateTime))
	fmt.Fprintf(cmd.OutOrStdout(), "TimeOnly %s \n", t.Format(time.TimeOnly))

}

func PrintNow() {
	now := time.Now()
	fmt.Printf("Now is %s \n", now.Format(time.RFC3339))
	fmt.Printf("Unix seconds %d \n", now.UnixMilli())
}

func PrintDate(t time.Time) {
	fmt.Printf("Date is %s \n", t.Format(time.RFC3339))
	fmt.Printf("Unix seconds %d \n", t.UnixMilli())
}

func PrintUnix(t time.Time, millis bool, seconds bool) {
	if millis {
		fmt.Printf("Unix Millis = %d \n", t.UnixMilli())
	}
	if seconds {
		fmt.Printf("Unix seconds = %d \n", t.Unix())
	}
}

func PrintTimeFormatted(t time.Time, format string) string {
	s := t.Format(format)
	return s
}

func TimeInfo() *models.TimeInfo {
	now := time.Now()
	name, offset := now.Zone()
	return &models.TimeInfo{
		Now:    now,
		Format: time.RFC850,
		TimeZoneInfo: &models.TimeZoneInfo{
			ZoneInfo: nil,
			Name:     name,
			Offset:   offset,
		},
	}
}

func Week() {
	for i := -1; i < 7; i++ {
		now := time.Now()
		now = now.AddDate(0, 0, i)
		fmt.Println(now)
	}

	now := time.Now()
	weekday := now.Weekday()
	switch weekday {
	case time.Sunday:
		now.AddDate(0, 0, -7)
	case time.Tuesday:
		now = now.AddDate(0, 0, -1)
	case time.Wednesday:
		now.AddDate(0, 0, -2)
	}

	fmt.Println(now)
}
