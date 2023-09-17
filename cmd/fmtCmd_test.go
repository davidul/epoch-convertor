package cmd

import (
	"bytes"
	"epc/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute_Fmt(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"fmt"})
	SetStaticTime(&pkg.MockTime{})
	RootCmd.Execute()
	assert.Equal(t, "ANSIC    Tue Jan  1 00:00:00 2019 \nKitchen  12:00AM \nRFC822   01 Jan 19 00:00 UTC \nRFC850   Tuesday, 01-Jan-19 00:00:00 UTC \nRFC1123  Tue, 01 Jan 2019 00:00:00 UTC \nRFC3339  2019-01-01T00:00:00Z \nUnixDate Tue Jan  1 00:00:00 UTC 2019 \nRubyDate Tue Jan 01 00:00:00 +0000 2019 \n", b.String())
}
