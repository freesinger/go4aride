package json

import (
	"log"
	"reflect"
	"runtime"
	"testing"
)

func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: mismatch, expect %v but %v", file, line, expected, actual)
		return false
	} else {
		log.Printf("expect %v got %v", expected, actual)
	}
	return true
}

var nestedArray string = `[
    [{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]}],
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
]`

var simpleLog string = `[[1],2,3,[4,5,6]]`

func TestParseArray(t *testing.T) {
	paths := []string{
		"$[0]",
		"$[2]",
		"$[3][1]",
		"$[0][1]",
		"$[0][0].last",
		"$[0][1].age",
		"$[0][1].nets[1]",
		"$[1].nets",
	}
	expected := []string{
		"[1]",
		"3",
		"5",
		"{\"first\": \"Roger\", \"last\": \"Craig\", \"age\": 68, \"nets\": [\"fb\", \"tw\"]}",
		"Murphy",
		"68",
		"\"tw\"",
		"[\"ig\", \"tw\"]",
	}
	for i := range paths {
		if i < 3 {
			actual := JsonExtract(simpleLog, paths[i]).String()
			ExpectEqual(t.Errorf, expected[i], actual)
		} else {
			actual := JsonExtract(nestedArray, paths[i]).String()
			ExpectEqual(t.Errorf, expected[i], actual)
		}
	}
}

var dbLog string = `{
  "t": {"$date": "2020-05-20T19:18:40.604+00:00"},
  "s": "I",
  "c": "NETWORK",
  "fav.movie": "Deer Hunter",
  "id": 51800,
  "f": [
    [{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]}],
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ],
  "ctx": "conn281",
  "msg": "client metadata",
  "attr": {
    "remote": "192.168.14.15:37666",
    "client": "conn281",
    "doc": {
      "application": {
        "name": "MongoDB Shell"
     },
      "driver": {
        "name": "MongoDB Internal Client",
        "version": "4.4.0"
     },
      "os": {
        "type": "Linux",
        "name": "CentOS Linux release 8.0.1905 (Core) ",
        "architecture": "x86_64",
        "version": "Kernel 4.18.0-80.11.2.el8_0.x86_64"
     }
   }
 },
  "tags": ["1", "2", "3"],
  "array": [[1],2,3,[4,5,6]]
}`

func TestParseObject(t *testing.T) {
	paths := []string{
		"$.t",
		"$.t.$date",
		"$.id",
		"$.fav\\.movie",
		"$.attr.remote",
		"$.attr.doc.os.architecture",
		"$.tags",
		"$.tags[2]",
		"$.array",
	}
	expected := []string{
		"{\"$date\": \"2020-05-20T19:18:40.604+00:00\"}",
		"2020-05-20T19:18:40.604+00:00",
		"51800",
		"Deer Hunter",
		"192.168.14.15:37666",
		"x86_64",
		"[\"1\", \"2\", \"3\"]",
		"\"3\"",
		"[[1],2,3,[4,5,6]]",
	}
	for i := range paths {
		actual := JsonExtract(dbLog, paths[i]).String()
		ExpectEqual(t.Errorf, expected[i], actual)
	}
}
