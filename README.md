# timeperiod
A GO (https://golang.org) package

Convert strings like "3m", "2w" and "5d" to seconds, milli-, micro- or nanoseconds

It returns an **int** value

Invalid strings will return 0 (zero)

## Units
`s` // seconds  
`m` // minutes  
`h` // hours  
`d` // days  
`w` // weeks  

## Examples
`timeperiod.Seconds("5m")` // 300  
`timeperiod.Milliseconds("5m")` // 300000  
`timeperiod.Seconds("2d")` // 172800  
`timeperiod.Nanoeconds("26w")` // 15724800000000000  
`timeperiod.Microconds("1h")` // 3600000000  
