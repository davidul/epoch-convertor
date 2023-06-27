package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Execute_Now(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"now"})
	RootCmd.Execute()
	fmt.Println("")
}
