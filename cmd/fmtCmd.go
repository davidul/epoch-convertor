package cmd

import (
	"epc/pkg"
	"fmt"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt with no arguments.",
	Long:  "fmt prints all predefined available time formats.",
	Short: "fmt print time formats.",
	Run: func(cmd *cobra.Command, args []string) {
		now := t.Now()
		pkg.PrintFormats(now, cmd)

		if reference, err := cmd.Flags().GetBool("ref"); err == nil && reference {
			fmt.Fprintf(cmd.OutOrStdout(), "============================================================================\n")
			fmt.Fprintf(cmd.OutOrStdout(), "Reference: https://golang.org/pkg/time/#pkg-constants\n")
			pkg.PrintReference(cmd)
		}
	}}

func init() {
	fmtCmd.Flags().Bool("ref", false, "include reference")
	RootCmd.AddCommand(fmtCmd)
}
