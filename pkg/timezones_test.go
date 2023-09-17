package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadZoneInfo(t *testing.T) {
	info := CountryCodes()
	assert.NotNil(t, info)
	assert.NotNil(t, info["CZ"])
}

func Test_ReadZoneInfoISO(t *testing.T) {
	iso := ReadZoneInfoISO()
	assert.NotNil(t, iso)
}

func Test_ZoneList(t *testing.T) {
	list := ZoneList(ReadZoneInfoISO())
	assert.NotNil(t, list)
}

func Test_TimeZoneList(t *testing.T) {
	list := TimeZoneList()
	assert.NotNil(t, list)
}

func Test_TimeZoneMap(t *testing.T) {
	zoneMap := TimeZoneMap()
	assert.NotNil(t, zoneMap)
}

func Test_TimeZoneMapByCountryCode(t *testing.T) {
	zoneMap := TimeZoneMapByCountryCode()
	assert.NotNil(t, zoneMap)
}

func Test_TimeZoneMapByOffset(t *testing.T) {
	zoneMap := TimeZoneMapByOffset()
	assert.NotNil(t, zoneMap)
}
