package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Execute(t *testing.T) {
	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"--help"})
	RootCmd.Execute()
	fmt.Println("")
}
