# timeperiod
A GO (https://golang.org) package

Convert strings like "3m", "2w" and "5d" to seconds, milli-, micro- or nanoseconds

It returns an **int** value

Invalid strings will return 0 (zero)

## Units
`ns` // nanoseconds  
`us` or `µs` // microseconds  
`ms` // milliseconds  
`s` // seconds  
`m` // minutes  
`h` // hours  
`d` // days  
`w` // weeks  

## Examples
`timeperiod.Seconds("26w")` // 15724800  
`timeperiod.Seconds("2d")` // 172800  
`timeperiod.Seconds("5m")` // 300  
`timeperiod.Milliseconds("5m")` // 300000  
`timeperiod.Microconds("1h")` // 3600000000  
`timeperiod.Nanoseconds("1us")` // 1000  
