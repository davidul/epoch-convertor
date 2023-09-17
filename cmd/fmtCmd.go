package cmd

import (
	"epc/pkg"
	"github.com/spf13/cobra"
	"time"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt with no arguments.",
	Long:  "fmt prints all predefined available time formats.",
	Short: "fmt print time formats.",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		pkg.PrintFormats(now)
	}}

func init() {
	RootCmd.AddCommand(fmtCmd)
}
