package cmd

import (
	"bytes"
	"epc/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var output = "Reference    01/01 12:00:00AM '19 +0000 \nANSIC    Tue Jan  1 00:00:00 2019 \nKitchen  12:00AM \nRFC822   01 Jan 19 00:00 UTC \nRFC822Z 01 Jan 19 00:00 +0000 \nRFC850   Tuesday, 01-Jan-19 00:00:00 UTC \nRFC1123  Tue, 01 Jan 2019 00:00:00 UTC \nRFC1123Z Tue, 01 Jan 2019 00:00:00 +0000 \nRFC3339  2019-01-01T00:00:00Z \nRFC3339Nano 2019-01-01T00:00:00Z \nUnixDate Tue Jan  1 00:00:00 UTC 2019 \nRubyDate Tue Jan 01 00:00:00 +0000 2019 \nStamp Jan  1 00:00:00 \nStamp Milli Jan  1 00:00:00.000 \nStamp Micro Jan  1 00:00:00.000000 \nStamp Nano Jan  1 00:00:00.000000000 \nDateOnly 2019-01-01 \nDateOnly 2019-01-01 00:00:00 \nTimeOnly 00:00:00 \n"

func TestExecute_Fmt(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"fmt"})
	SetStaticTime(&pkg.MockTime{})
	RootCmd.Execute()
	assert.Equal(t, output, b.String())
}
