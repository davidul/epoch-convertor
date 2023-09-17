package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "now",
	Long:  "now displays current time in unix epoch",
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultFormat := viper.GetString("default-format")
		if defaultFormat == "" {
			defaultFormat = "2006-01-02 15:04:05"
		}
		now := t.Now()
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "Current date and time is: %s\n",
			now.Format(defaultFormat))
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch seconds: %d\n", now.Unix())
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch miliseconds: %d\n", now.UnixMilli())
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Unix epoch microseconds: %d\n", now.UnixMicro())
		if err != nil {
			return err
		}

		name, offset := now.Zone()
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Local Timezone: %s\n", name)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(cmd.OutOrStdout(), "Offset: %d\n", offset)
		if err != nil {
			return err
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(nowCmd)
}
