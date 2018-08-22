# timeperiod
A GO (https://golang.org) package

Convert strings like "2w" and "3d" to seconds, milli-, micro- or nanoseconds

It returns an **int** value

Invalid strings will return 0 (zero)

## Examples
`timeperiod.Seconds("5m")` // 300  
`timeperiod.Milliseconds("5m")` // 300000  
`timeperiod.Seconds("2d")` // 172800  
