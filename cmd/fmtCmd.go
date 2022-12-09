package cmd

import (
	"epoch-convertor/pkg"
	"github.com/spf13/cobra"
	"time"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Long:  "fmt",
	Short: "fmt",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		pkg.PrintFormats(now)
	}}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
