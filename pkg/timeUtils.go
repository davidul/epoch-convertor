package pkg

import (
	"fmt"
	"time"
)

func PrintFormats(t time.Time) {
	fmt.Printf("ANSIC    %s \n", t.Format(time.ANSIC))
	fmt.Printf("Kitchen  %s \n", t.Format(time.Kitchen))
	fmt.Printf("RFC822   %s \n", t.Format(time.RFC822))
	fmt.Printf("RFC850   %s \n", t.Format(time.RFC850))
	fmt.Printf("RFC1123  %s \n", t.Format(time.RFC1123))
	fmt.Printf("RFC3339  %s \n", t.Format(time.RFC3339))
	fmt.Printf("UnixDate %s \n", t.Format(time.UnixDate))
	fmt.Printf("RubyDate %s \n", t.Format(time.RubyDate))
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
