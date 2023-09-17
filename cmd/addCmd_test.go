package cmd

import (
	"bytes"
	"epc/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Execute_Add(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"add", "--day", "1"})
	SetStaticTime(&pkg.MockTime{})
	err := RootCmd.Execute()
	if err != nil {
		return
	}
	assert.Equal(t, "Now Tue Jan  1 00:00:00 UTC 2019\nAdd 0 year(s) 0 month(s) 1 day(s)\nWed Jan  2 00:00:00 UTC 2019\n\t Unix epoch seconds: 1546387200\n\t Unix epoch miliseconds: 1546300800000\n", b.String())
}
