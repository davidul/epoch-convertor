package pkg

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ZoneInfo struct {
	CountryCode string
	coordinates string
	TimeZone    string
	comments    string
}

func (z ZoneInfo) String() string {
	s := "countryCode: " + z.CountryCode + "\n"
	s = s + "timeZone: " + z.TimeZone
	return s
}

func CountryCodes() map[string]string {

	tzIsoFile, err := ioutil.ReadFile("/usr/share/zoneinfo/iso3166.tab")
	if err != nil {
		fmt.Println(err)
	}

	tzMap := make(map[string]string)

	lines := strings.Split(string(tzIsoFile), "\n")
	for i := range lines {
		line := lines[i]
		if index := strings.Index(line, "#"); index > -1 {
			line = line[:index]
		}

		trimedLine := strings.TrimSpace(line)
		if trimedLine == "" {
			continue
		}
		tz := strings.Split(trimedLine, "\t")
		tzMap[tz[0]] = tz[1]

	}

	return tzMap
}

// ReadZoneInfoISO read zone.tab file
// returns map, key is time zone
func ReadZoneInfoISO() map[string]ZoneInfo {
	file, err := ioutil.ReadFile("/usr/share/zoneinfo/zone.tab")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	tzMap := make(map[string]ZoneInfo)

	lines := strings.Split(string(file), "\n")

	for i := range lines {
		line := lines[i]
		if index := strings.Index(line, "#"); index > -1 {
			line = line[:index]
		}

		trimedLine := strings.TrimSpace(line)
		if trimedLine == "" {
			continue
		}

		split := strings.Split(trimedLine, "\t")
		var comment string
		if len(split) == 4 {
			comment = split[3]
		} else {
			comment = ""
		}
		info := ZoneInfo{
			CountryCode: split[0],
			coordinates: split[1],
			TimeZone:    split[2],
			comments:    comment,
		}

		tzMap[split[2]] = info

	}

	return tzMap
}

func ZoneList(zoneMap map[string]ZoneInfo) []string {

	i := make([]string, len(zoneMap))
	cnt := 0
	for k, _ := range zoneMap {
		i[cnt] = k
		cnt++
	}

	return i
}
