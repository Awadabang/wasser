# wasser

Convert Struct to Map in Go.

``` go
type Server struct {
	Name  string `json:"name"`
	Age   int64  `json:"age"`
}

s := `{"name": "alex"}`
var server Server
m, err := Struct2Map([]byte(s), &server)
if err != nil {
	log.Fatalln(err)
}
log.Println(m)
// output: map[name:alex]
```
