# tt
a simple time tracker for go

# Example

```go
// GetDataStatic Return static data
func GetDataStatic() []Load {
	defer tt.Track(time.Now(), "GetDataStaticLoads")
	f, _ := fs.New()
	file, _ := f.Open("/loads.gob")
	dec := gob.NewDecoder(file)
	v := []Load{}
	dec.Decode(&v)
	return v
}
```
output: 
```
GetDataStaticLoads took 75.563159ms
```

