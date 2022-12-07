package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "epc",
	Short: "epc",
	Long:  "epc",
	Run: func(cmd *cobra.Command, args []string) {
		epochInt, err := cmd.Flags().GetInt64("epoch")
		if err != nil {
			fmt.Println(err)
		}
		unix := time.Unix(epochInt, 0)
		milli := time.UnixMilli(epochInt)
		fmt.Printf("%s \n", unix.Format(time.RFC850))
		fmt.Printf("%s \n", milli.Format(time.RFC850))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.Flags().Int64("epoch", 0, "seconds since epoch")
}
