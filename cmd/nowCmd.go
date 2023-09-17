package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "now",
	Long:  "now displays current time in unix epoch",
	RunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now()
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "Current date and time is: %s \n",
			now.Format(viper.GetString("default-format")))
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch seconds: %d \n", now.Unix())
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch miliseconds: %d \n", now.UnixMilli())
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch microseconds: %d \n", now.UnixMicro())
		if err != nil {
			return err
		}

		t2 := time.Now()
		name, offset := t2.Zone()
		fmt.Fprintf(cmd.OutOrStdout(), "Local Timezone: %s\n", name)
		fmt.Fprintf(cmd.OutOrStdout(), "Offset: %d\n", offset)
		return err
	},
}

func init() {
	RootCmd.AddCommand(nowCmd)
}
