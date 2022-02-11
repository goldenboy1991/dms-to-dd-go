# dms-to-dd-go

Сonverts the DMS(degrees,minutes,seconds) format to DD(decimal degrees).

## Getting started

Import the module:
```
dmstodd "github.com/goldenboy1991/dms-to-dd-go"
```
Create a DMS object using the NewDMS method, pass values into it in the DMS format (degrees, minutes, seconds):
```
dms := dmstodd.NewDMS('44/1','44/1','39/1')
```
Convert coordinates from DMS to DD:

```
dd, err := dmsLongitude.ConverterDMSToDD()
if err != nil {
    log.println(err)
    return
}
log.println('dd:', dd) // 44.744166666666665
```

## Authors and acknowledgment
goldenboy1991

## License
MIT
