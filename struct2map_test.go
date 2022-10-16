package wasser

import (
	"log"
	"testing"
)

type Grade struct {
	Math    int64 `json:"math"`
	English int64 `json:"english"`
}

type Server struct {
	Name  string `structs:"name"`
	Age   int64  `structs:"age"`
	Grade *Grade `structs:"grade"`
}

func TestStruct2Map(t *testing.T) {
	s := `{"name": "alex", "grade": {"english":200}}`
	var server Server
	m, err := Struct2map([]byte(s), &server)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(m)
}
