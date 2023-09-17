package models

import "time"

type ZoneInfo struct {
	CountryCode string
	Coordinates string
	TimeZone    string
	Comments    string
}

type TimeZoneInfo struct {
	ZoneInfo *ZoneInfo
	Name     string
	Offset   int
}

type TimeInfo struct {
	Now          time.Time
	Format       string
	TimeZoneInfo *TimeZoneInfo
}

func (z ZoneInfo) String() string {
	s := "countryCode: " + z.CountryCode + "\n"
	s = s + "timeZone: " + z.TimeZone
	return s
}
