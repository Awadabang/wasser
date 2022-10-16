# wasser

Convert Struct to Map in Go.

``` go
type Server struct {
	Name  string `structs:"name"`
	Age   int64  `structs:"age"`
}

s := `{"name": "alex"}`
var server Server
m, err := Struct2map([]byte(s), &server)
if err != nil {
	log.Fatalln(err)
}
log.Println(m)
// output: map[name:alex]
```
