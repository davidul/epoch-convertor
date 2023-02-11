package pkg

import (
	"fmt"
	"github.com/davidul/go-vic/linkedlist"
	"io/ioutil"
	"sort"
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

func ReadZoneInfoISO() map[string]linkedlist.LinkedList[ZoneInfo] {
	file, err := ioutil.ReadFile("/usr/share/zoneinfo/zone.tab")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	tzMap := make(map[string]linkedlist.LinkedList[ZoneInfo])

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

		list, exists := tzMap[split[0]]
		if exists {
			list.Add(info)
		} else {
			l := linkedlist.LinkedList[ZoneInfo]{}
			l.Add(info)
			tzMap[split[0]] = l
		}

	}

	return tzMap
}

func ZoneList(zoneMap map[string]linkedlist.LinkedList[ZoneInfo]) linkedlist.LinkedList[ZoneInfo] {

	l := linkedlist.LinkedList[ZoneInfo]{}
	for _, v := range zoneMap {
		l.AddAll(&v)
	}

	l.ToArray()
	return l
}

func ZoneListSortedArray(zoneMap map[string]linkedlist.LinkedList[ZoneInfo]) []string {
	list := ZoneList(zoneMap)
	array := list.ToArray()
	sarray := make([]string, len(array))
	for i, a := range array {
		sarray[i] = a.TimeZone
	}
	sort.Strings(sarray)
	return sarray
}
