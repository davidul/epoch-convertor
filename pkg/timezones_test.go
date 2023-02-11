package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadZoneInfo(t *testing.T) {
	info := CountryCodes()
	assert.NotNil(t, info)
	assert.NotNil(t, info["CZ"])
}

func TestReadZoneInfoISO(t *testing.T) {
	iso := ReadZoneInfoISO()
	assert.NotNil(t, iso)
}
