package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var parseCmd = &cobra.Command{
	Use:   "parse date/time in different formats",
	Long:  "parse",
	Short: "parse.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read date from flag or first positional arg
		date, err := cmd.Flags().GetString("date")
		if err != nil {
			return fmt.Errorf("failed to read --date flag: %w", err)
		}
		if date == "" && len(args) > 0 {
			date = args[0]
		}
		if date == "" {
			return fmt.Errorf("date is required (use --date or provide as first arg)")
		}

		// input layout
		format, err := cmd.Flags().GetString("in")
		if err != nil {
			return fmt.Errorf("failed to read --in flag: %w", err)
		}
		if format == "" {
			mapString := viper.GetStringMapString("date-formats")
			format = mapString["default-format"]
		}
		if format == "" {
			return fmt.Errorf("input format is required (use --in or set date-formats.default-format in config)")
		}

		// timezone / location
		tz, err := cmd.Flags().GetString("tz")
		if err != nil {
			return fmt.Errorf("failed to read --tz flag: %w", err)
		}
		var loc *time.Location
		if tz != "" {
			l, err := time.LoadLocation(tz)
			if err != nil {
				return fmt.Errorf("failed to load timezone %s: %w", tz, err)
			}
			loc = l
		} else {
			loc = time.Local
		}

		// parse using location so layouts without zone are interpreted in the requested tz
		parsed, err := time.ParseInLocation(format, date, loc)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}

		// output handling
		out, err := cmd.Flags().GetString("out")
		if err != nil {
			return fmt.Errorf("failed to read --out flag: %w", err)
		}

		switch out {
		case "unix":
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), parsed.Unix())
		case "rfc3339":
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), parsed.Format(time.RFC3339))
		case "":
			// default behavior preserves previous output: human stamp + unix seconds
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), parsed.Format(time.Stamp))
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), parsed.Unix())
		default:
			// treat out as a Go time layout
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), parsed.Format(out))
		}

		return nil
	},
}

func init() {
	parseCmd.Flags().String("in", "", "input layout (Go time layout). If empty the config value date-formats.default-format is used")
	parseCmd.Flags().String("date", "", "date to parse (or provide as first arg)")
	parseCmd.Flags().String("out", "", "output. 'unix' prints unix seconds, 'rfc3339' prints RFC3339, or provide a Go time layout")
	parseCmd.Flags().String("tz", "", "timezone IANA string (e.g. Europe/London). If empty uses local timezone")
	RootCmd.AddCommand(parseCmd)
}
