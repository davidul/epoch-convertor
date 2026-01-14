package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	format := "2006-01-02T15:04:05"
	date := "2006-01-02T15:04:05"
	parse, err := time.Parse(format, date)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(parse.Unix())
}
