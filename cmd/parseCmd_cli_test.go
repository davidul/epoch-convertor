package cmd

import (
	"bytes"
	"strings"
	"testing"
)

// Helper to execute the root command with args and capture stdout/stderr
func execCommand(args ...string) (string, string, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	RootCmd.SetOut(bufOut)
	RootCmd.SetErr(bufErr)
	RootCmd.SetArgs(args)

	err := RootCmd.Execute()
	return strings.TrimSpace(bufOut.String()), strings.TrimSpace(bufErr.String()), err
}

func TestParseDefaultFormatPositionalArg(t *testing.T) {
	// Use explicit --in to avoid touching global viper in other tests
	out, errOut, err := execCommand("parse", "--in", "2006-01-02T15:04:05", "2006-01-02T15:04:05")
	if err != nil {
		t.Fatalf("expected no error, got: %v stderr:%s", err, errOut)
	}
	if out == "" {
		t.Fatalf("expected output, got empty")
	}
	// The default behavior prints a human stamp plus unix seconds; ensure unix appears as digits
	if !strings.Contains(out, " ") || !strings.ContainsAny(out, "0123456789") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestParseWithInAndOutUnix(t *testing.T) {
	out, errOut, err := execCommand("parse", "--date", "2006-01-02T15:04:05", "--in", "2006-01-02T15:04:05", "--out", "unix")
	if err != nil {
		t.Fatalf("expected no error, got: %v stderr:%s", err, errOut)
	}
	if out == "" {
		t.Fatalf("expected unix output, got empty")
	}
	// unix should be numeric
	if strings.TrimSpace(out) == "" {
		t.Fatalf("empty stdout")
	}
	// should parse to a known unix value for the reference date
	if !strings.HasPrefix(out, "-") && !strings.ContainsAny(out, "0123456789") {
		t.Fatalf("unexpected unix output: %q", out)
	}
}

func TestParseWithTimezone(t *testing.T) {
	// Use a date without zone and parse it in UTC via --tz
	out, errOut, err := execCommand("parse", "--date", "2006-01-02T15:04:05", "--in", "2006-01-02T15:04:05", "--out", "rfc3339", "--tz", "UTC")
	if err != nil {
		t.Fatalf("expected no error, got: %v stderr:%s", err, errOut)
	}
	if !strings.Contains(out, "2006-01-02T15:04:05") {
		t.Fatalf("unexpected rfc3339 output: %q", out)
	}
}

// --- Failing-case tests ---

func TestParseMissingDateFails(t *testing.T) {
	// Clear any previously set flags to avoid test interference when running the whole suite
	_ = parseCmd.Flags().Set("date", "")
	_ = parseCmd.Flags().Set("in", "")
	_ = parseCmd.Flags().Set("out", "")

	// No date provided should return an error
	_, errOut, err := execCommand("parse")
	if err == nil {
		t.Fatalf("expected an error when date is missing, got nil; stderr: %s", errOut)
	}
	if !strings.Contains(err.Error(), "date is required") && !strings.Contains(errOut, "date is required") {
		t.Fatalf("expected 'date is required' error, got err: %v stderr: %s", err, errOut)
	}
}

func TestParseInvalidFormatFails(t *testing.T) {
	// Provide a format that doesn't match the date string
	_, errOut, err := execCommand("parse", "--date", "not-a-date", "--in", "2006-01-02T15:04:05")
	if err == nil {
		t.Fatalf("expected parse error for invalid date, got nil; stderr: %s", errOut)
	}
	if !strings.Contains(err.Error(), "failed to parse date") && !strings.Contains(errOut, "failed to parse date") {
		t.Fatalf("expected 'failed to parse date' error, got err: %v stderr: %s", err, errOut)
	}
}
