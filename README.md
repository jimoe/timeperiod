# timeperiod

A [GO](https://golang.org) module  
See [pkg.go.dev](https://pkg.go.dev/github.com/jimoe/timeperiod) for documentation

Convert strings like "3m", "2w" and "5d".

It can convert to:
- either **int** values for seconds, milli-, micro- or nanoseconds
- or [time.Duration](https://pkg.go.dev/time#Duration)

Invalid strings will return 0 (zero)

If no unit is given it will assume that the value is in seconds

## Version 2
Version 2 uses generics to allow a custom type `timeperiod.Period` (`type Period string`) 
or string as you prefer. This is to clarify that the string is a time period and not just 
a string when used in configurations, as a function parameter or similar.

Due to generics it now requires to use go 1.18 or later.

## Units
`ns` // nanoseconds  
`us` // microseconds  
`µs` // microseconds  
`ms` // milliseconds  
`s` // seconds  
`m` // minutes  
`h` // hours  
`d` // days  
`w` // weeks  

## Examples
`timeperiod.Duration("5ns")` // 5  
`timeperiod.Seconds("26w")` // 15724800  
`timeperiod.Seconds("2d")` // 172800  
`timeperiod.Seconds("5m")` // 300  
`timeperiod.Milliseconds("5m")` // 300000  
`timeperiod.Microconds("1h")` // 3600000000  
`timeperiod.Nanoseconds("1us")` // 1000  
`timeperiod.Nanoseconds("1µs")` // 1000  
