package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Long:  "fmt",
	Short: "fmt",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		fmt.Printf("ANSIC    %s \n", now.Format(time.ANSIC))
		fmt.Printf("Kitchen  %s \n", now.Format(time.Kitchen))
		fmt.Printf("RFC822   %s \n", now.Format(time.RFC822))
		fmt.Printf("RFC850   %s \n", now.Format(time.RFC850))
		fmt.Printf("RFC1123  %s \n", now.Format(time.RFC1123))
		fmt.Printf("RFC3339  %s \n", now.Format(time.RFC3339))
		fmt.Printf("UnixDate %s \n", now.Format(time.UnixDate))
		fmt.Printf("RubyDate %s \n", now.Format(time.RubyDate))
	}}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
