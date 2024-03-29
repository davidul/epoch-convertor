package pkg

import (
	"epc/models"
	"epc/pkg/csv"
	"time"
)

//Sample of iso3166.tab, comments start with '#' character
//#country-
//#code	name of country, territory, area, or subdivision
//AD	Andorra
//AE	United Arab Emirates
//AF	Afghanistan
//AG	Antigua & Barbuda

func CountryCodes() map[string]string {

	parser := csv.NewParser('\t', '#', nil)
	tzIsoFile := parser.ParseFile(csv.ReadFileAsByteBuffer("/usr/share/zoneinfo/iso3166.tab"))
	//tzIsoFile, err := ioutil.ReadFile("/usr/share/zoneinfo/iso3166.tab")
	//if err != nil {
	//	fmt.Println(err)
	//}

	tzMap := make(map[string]string)
	for i := range tzIsoFile.Records {
		record := tzIsoFile.Records[i]
		tzMap[record.Value[0]] = record.Value[1]
	}

	return tzMap
}

// ReadZoneInfoISO read zone.tab file
// returns map, key is time zone
// Country code sample
//
// # This table is intended as an aid for users, to help them select timezones
// # appropriate for their practical needs.  It is not intended to take or
// # endorse any position on legal or territorial claims.
// #
// #country-
// #code	coordinates	TZ			comments
// AD	+4230+00131	Europe/Andorra
// AE	+2518+05518	Asia/Dubai
// AF	+3431+06912	Asia/Kabul
// AG	+1703-06148	America/Antigua
// AI	+1812-06304	America/Anguilla
// AL	+4120+01950	Europe/Tirane
// AM	+4011+04430	Asia/Yerevan
// AO	-0848+01314	Africa/Luanda
// AQ	-7750+16636	Antarctica/McMurdo	New Zealand time - McMurdo, South Pole
func ReadZoneInfoISO() map[string]*models.ZoneInfo {

	parser := csv.NewParser('\t', '#', nil)
	file := parser.ParseFile(csv.ReadFileAsByteBuffer("/usr/share/zoneinfo/zone.tab"))

	tzMap := make(map[string]*models.ZoneInfo)

	for i := range file.Records {
		record := file.Records[i]
		info := &models.ZoneInfo{
			CountryCode: record.Value[0],
			Coordinates: record.Value[1],
			TimeZone:    record.Value[2],
			Comments:    "",
		}

		tzMap[record.Value[2]] = info

	}

	return tzMap
}

// ZoneList returns a list of time zone
func ZoneList(zoneMap map[string]*models.ZoneInfo) []string {
	i := make([]string, len(zoneMap))
	cnt := 0
	for k, _ := range zoneMap {
		i[cnt] = k
		cnt++
	}

	return i
}

func TimeZoneList() []*models.TimeZoneInfo {
	iso := ReadZoneInfoISO()
	infos := make([]*models.TimeZoneInfo, 0, len(iso))
	for _, z := range iso {
		l, _ := time.LoadLocation(z.TimeZone)
		name, offset := time.Now().In(l).Zone()
		infos = append(infos, &models.TimeZoneInfo{
			ZoneInfo: z,
			Name:     name,
			Offset:   offset,
		})
	}

	return infos
}

func TimeZoneMap() map[string]*models.TimeZoneInfo {
	iso := ReadZoneInfoISO()
	infos := make(map[string]*models.TimeZoneInfo, len(iso))
	for key, z := range iso {
		l, _ := time.LoadLocation(z.TimeZone)
		name, offset := time.Now().In(l).Zone()
		infos[key] = &models.TimeZoneInfo{
			ZoneInfo: z,
			Name:     name,
			Offset:   offset,
		}
	}

	return infos
}

func TimeZoneMapByCountryCode() map[string][]*models.TimeZoneInfo {
	iso := ReadZoneInfoISO()
	infos := make(map[string][]*models.TimeZoneInfo, len(iso))
	for _, z := range iso {
		l, _ := time.LoadLocation(z.TimeZone)
		name, offset := time.Now().In(l).Zone()
		infos[z.CountryCode] = append(infos[z.CountryCode], &models.TimeZoneInfo{
			ZoneInfo: z,
			Name:     name,
			Offset:   offset,
		})
	}

	return infos
}

func TimeZoneMapByCountryCodeAndTimeZone() map[string]map[string]*models.TimeZoneInfo {
	iso := ReadZoneInfoISO()
	infos := make(map[string]map[string]*models.TimeZoneInfo, len(iso))
	for _, z := range iso {
		l, _ := time.LoadLocation(z.TimeZone)
		name, offset := time.Now().In(l).Zone()
		if _, ok := infos[z.CountryCode]; !ok {
			infos[z.CountryCode] = make(map[string]*models.TimeZoneInfo)
		}
		infos[z.CountryCode][z.TimeZone] = &models.TimeZoneInfo{
			ZoneInfo: z,
			Name:     name,
			Offset:   offset,
		}
	}

	return infos
}

func TimeZoneMapByOffset() map[int][]*models.TimeZoneInfo {
	iso := ReadZoneInfoISO()
	infos := make(map[int][]*models.TimeZoneInfo, len(iso))
	for _, z := range iso {
		l, _ := time.LoadLocation(z.TimeZone)
		name, offset := time.Now().In(l).Zone()
		timeZoneInfo := &models.TimeZoneInfo{
			ZoneInfo: z,
			Name:     name,
			Offset:   offset,
		}

		if val, ok := infos[offset]; ok {
			infos[offset] = append(val, timeZoneInfo)
		} else {
			zoneInfos := make([]*models.TimeZoneInfo, 1, 2)
			zoneInfos[0] = timeZoneInfo
			infos[offset] = zoneInfos
		}

	}

	return infos
}
