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
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Grade *Grade `json:"grade"`
}

func TestStruct2Map(t *testing.T) {
	s := `{"name": "alex", "grade": {"english":200}}`
	var server Server
	m, err := Struct2Map([]byte(s), &server)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(m)
}
