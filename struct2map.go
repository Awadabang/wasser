package wasser

import (
	"errors"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
)

func Struct2Map(b []byte, s any) (map[string]any, error) {
	result := make(map[string]any)
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Pointer || v.IsNil() {
		return nil, errors.New("s is not a pointer or s is nil")
	}

	// get the underlying element
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("s is not a struct")
	}

	// get type
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("json")

		// don't check if it's omitted
		if tag == "-" {
			continue
		}

		res := strings.Split(tag, ",")
		jsonKey := res[0]

		jsonVal := gjson.GetBytes(b, jsonKey)
		if jsonType := jsonVal.Type; jsonType == gjson.Null {
			continue
		}

		result[jsonKey] = jsonVal.Value()
	}

	return result, nil
}
