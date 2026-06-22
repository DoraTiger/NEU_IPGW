package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

var OutputJSON bool

func Output(data interface{}, textFormatter func()) {
	if OutputJSON {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "JSON marshal error: %v\n", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		textFormatter()
	}
}

func ToJSON(data interface{}) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "%v"}`, err)
	}
	return string(jsonData)
}

func MapToJSON(m map[string]interface{}) string {
	return ToJSON(m)
}

func StructToJSON(s interface{}) string {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ToJSON(s)
	}

	m := make(map[string]interface{})
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if field.PkgPath == "" {
			m[field.Name] = value.Interface()
		}
	}
	return ToJSON(m)
}