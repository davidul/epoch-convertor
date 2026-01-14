package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountryCodes(t *testing.T) {
	codes := CountryCodes()
	if len(codes) == 0 {
		t.Error("CountryCodes() returned an empty array")
	}

	assert.True(t, codes["US"] == "United States")
}
