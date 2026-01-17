# Epoch converter
Simple util for time formatting

## Commands

Without parameters prints current datetime and its 
representation in unix seconds.

```shell
./epc
Time in custom format 2022-12-09T15:03:14+01:00
Unix seconds = 1670594594
```

Available commands:

```shell
    fmt         prints current time in different formats
    add         adds time to current time
    help        Help about any command
    now         prints current time
    parse       converts unix time stamp to date
    stopwatch   runs stop watch in terminal
    tz          tz
    unix        unix converts unix timestamp with different precisions
  ```

Available flags

```
      --day int         +/-day
      --month int       +/-month
      --year int        +/-year
      --format string   format of date time (default "2006-01-02T15:04:05Z07:00")
  -h, --help            help for epc
      --millis          unix time in millis
      --seconds         unix time in seconds (default true)
```

For example tomorrow (+1 day)

```shell
./epc --day 1
```

Six months ago

```shell
./epc --month -6
```

Millis one year ago
```shell
./epc --year -1 --millis
```

Enable disable seconds/milliseconds with
`--seconds` and `--millis`

```shell
./epc --millis --seconds=false
Time in custom format 2022-12-09T15:12:57+01:00
Unix Millis = 1670595177743 
```

Move time there and back with `--day`, `--month` and
`--year`. Use +/- to add/subtract from current date.

```shell
./epc --day 100 --month -3 --year 10                                                           
Time in custom format 2032-12-18T15:41:21+01:00
Unix seconds = 1986993681 
```

## now

`now` just prints current data with unix timestamps.

```shell
./epc now
Current date and time is: 2023-09-27T19:01:34
Unix epoch seconds: 1695834094
Unix epoch miliseconds: 1695834094048
Unix epoch microseconds: 1695834094048884
Local Timezone: CEST
Offset: 7200
```

## stopwatch

`stopwatch` runs stop watch in terminal
```shell
./epc stopwatch 
00:00:05
```

## unix

`unix` converts unix time stamp to date, you can 
provide the precision of the timestamp.

```shell
./epc unix 1234556666 --millis
```

```shell
--micros          unix time in microseconds
--millis          unix time in millis
--seconds         unix time in seconds
```
If no precision is provided, it will try all.
```shell
./epc parse 1234556666 
From millis Thursday, 15-Jan-70 07:55:56 CET 
From micros Thursday, 01-Jan-70 01:20:34 CET 
From seconds Friday, 13-Feb-09 21:24:26 CET 

```

## fmt

fmt - prints current time in different formats

```shell
./epc fmt

ANSIC    Sun Nov  5 17:09:13 2023 
Kitchen  5:09PM 
RFC822   05 Nov 23 17:09 CET 
RFC822Z 05 Nov 23 17:09 +0100 
RFC850   Sunday, 05-Nov-23 17:09:13 CET 
RFC1123  Sun, 05 Nov 2023 17:09:13 CET 
RFC1123Z Sun, 05 Nov 2023 17:09:13 +0100 
RFC3339  2023-11-05T17:09:13+01:00 
RFC3339Nano 2023-11-05T17:09:13.083371+01:00 
UnixDate Sun Nov  5 17:09:13 CET 2023 
RubyDate Sun Nov 05 17:09:13 +0100 2023 
Stamp Nov  5 17:09:13 
Stamp Milli Nov  5 17:09:13.083 
Stamp Micro Nov  5 17:09:13.083371 
Stamp Nano Nov  5 17:09:13.083371000 
DateOnly 2023-11-05 
DateOnly 2023-11-05 17:09:13 
TimeOnly 17:09:13 
```

You can include the reference format itself
```shell
./epc fmt --ref
```

# parse

```shell
./epc parse --date 2006-01-03T15:04:05
./epc parse 2006-01-02T15:04:05
./epc parse --date "2023-09-15T14:30:00" --in "2006-01-02T15:04:05"

./epc parse --date "2023-09-15T14:30:00" --in "2006-01-02T15:04:05" --out unix

./epc parse --date "2023-09-15T14:30:00" --in "2006-01-02T15:04:05" --tz "Europe/London" --out rfc3339

./epc parse --date "2023-09-15T14:30:00" --in "2006-01-02T15:04:05" --out "2006-01-02 15:04:05"



```

## tz
