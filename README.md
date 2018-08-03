# tt
A simple time tracker for go in arround 40 lines of code.  
Allows you to quickly check how much time a function takes to run.

## Doc
https://godoc.org/github.com/chneau/tt

## Install
```bash
go get github.com/chneau/tt
```

## Usage

Simply add `defer tt.T()()` to a function to know how much times the function takes to finish (from the line you have written `defer tt.T()()`).

## Example
###### Code
```go
// GetDataStatic Return static data
func GetDataStatic() []Load {
	defer tt.Track(time.Now(), "GetDataStaticLoads") // This
	defer tt.T()()                                   // This one is a shorthand
	f, _ := fs.New()
	file, _ := f.Open("/loads.gob")
	dec := gob.NewDecoder(file)
	v := []Load{}
	dec.Decode(&v)
	return v
}
```
###### Output
```
[TRACK] 2018/08/03 12:12:03 <[main.GetDataStatic]> 75.563159ms
[TRACK] 2018/08/03 12:12:03 <[GetDataStaticLoads]> 75.563159ms
```

