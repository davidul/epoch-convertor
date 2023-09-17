package cmd

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_Execute_Add(t *testing.T) {
	now := func() time.Time {
		return time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	now()

	b := new(bytes.Buffer)
	RootCmd.SetOut(b)
	RootCmd.SetErr(b)
	RootCmd.SetArgs([]string{"add", "--day", "1"})
	RootCmd.Execute()
	fmt.Println(b.String())
}
