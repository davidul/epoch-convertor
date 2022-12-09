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

Enable disable seconds/milliseconds with
`--seconds` and `--millis`

```shell
./epc --millis --seconds=false                                                                 ✔  base   15:08:39  
Time in custom format 2022-12-09T15:12:57+01:00
Unix Millis = 1670595177743 
```

Move time there and back with `--day`, `--month` and
`--year`. Use +/- to add/subtract from current date.

```shell
./epc --day 100 --month -3 --year 10                                                           ✔  base   15:12:57  
Time in custom format 2032-12-18T15:41:21+01:00
Unix seconds = 1986993681 
```


fmt - prints current time in different formats

```shell
./epoch-converter fmt
ANSIC    Wed Dec  7 06:22:00 2022 
Kitchen  6:22AM 
RFC822   07 Dec 22 06:22 CET 
RFC850   Wednesday, 07-Dec-22 06:22:00 CET 
RFC1123  Wed, 07 Dec 2022 06:22:00 CET 
RFC3339  2022-12-07T06:22:00+01:00 
UnixDate Wed Dec  7 06:22:00 CET 2022 
RubyDate Wed Dec 07 06:22:00 +0100 2022 
```

