package cmd

import (
	"bytes"
	"epc/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Execute_Now(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"now"})
	SetStaticTime(&pkg.MockTime{})
	RootCmd.Execute()
	assert.Equal(t, "Current date and time is: 2019-01-01 00:00:00\nUnix epoch seconds: 1546300800\nUnix epoch miliseconds: 1546300800000\nUnix epoch microseconds: 1546300800000000\nLocal Timezone: UTC\nOffset: 0\n", b.String())
}
