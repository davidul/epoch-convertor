package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "now",
	Long:  "now displays current time in unix epoch",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		fmt.Printf("%s \n", now.Format("2006-01-02T15:04:05"))
		fmt.Printf("Unix epoch seconds: %d \n", now.Unix())
		fmt.Printf("Unix epoch miliseconds: %d \n", now.UnixMilli())
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
