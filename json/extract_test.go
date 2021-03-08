package json

import (
	"github.com/freesinger/go4aride/utils"
	"testing"
)
const dbLog = `{
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

func TestParseKeys(t *testing.T) {
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
	results := []string{
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
		utils.ExpectEqual(results[i], Extract(dbLog, paths[i]).String())
	}
}

const nestedArray = `[
    [{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]}],
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}]`
const simpleLog = `[[1],2,3,[4,5,6]]`
func TestParseArray(t *testing.T) {
	paths := []string{
		"$[0]",
		"$[2]",
		"$[3][1]",
		"$[]",
	}
	results := []string{
		"[1]",
		"3",
		"5",
	}
	for i := range paths {
		if i < 3 {
			utils.ExpectEqual(results[i], Extract(simpleLog, paths[i]).String())
		} else {

		}
	}
}

func TestUnescapeString(t *testing.T) {
	path := `{"tags": ["1", "2", "3"]}`
	utils.ExpectEqual("\"3\"", Extract(path, "$.tags[2]").String())
}
