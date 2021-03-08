package main

import (
	"github.com/freesinger/go4aride/json"
	"github.com/freesinger/go4aride/utils"
)

func main() {
	const jsonDoc = `[
    [{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]}],
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`
	const jsonDoc1 = `{
  "t": {
    "$date": "2020-05-20T19:18:40.604+00:00"
 },
  "1": "number",
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
	const json3  = `{"array": [[1],2,3,[4,5,6]]}`
	//fmt.Println(strings.Split(".a.v.c.x[1]", "."))
	//res := json.Extract(jsonDoc, "$[0][1].last")
	//res := json.Extract(jsonDoc, "$.0.1.last")
	//utils.ExpectEqual("Craig", res.String())
	res2 := json.Extract(jsonDoc1, "$.tags[0]")
	utils.ExpectEqual("1", res2.String())
	r := json.Extract(json3, "$.array[3][2]")
	utils.ExpectEqual("6", r.String())
	res1 := json.Extract(jsonDoc, "$[0][1].last")
	utils.ExpectEqual("Craig", res1.String())

	paths := []string{"$.fav\\.movie", "$.f[0][1].last","$.f[1].last", "$.f[2].nets[0", "$.attr.doc.driver.name", "$.array.3.2"}
	expected := []string{"Deer Hunter", "Craig", "Murphy", "","MongoDB Internal Client", "6"}
	//results := json.ExtractMany(jsonDoc, paths...)
	//utils.ExpectEqual(len(results), len(paths))
	// log.Printf("%v\n", gjson.ExtractMany(json, paths...))
	for i, path := range paths {
		result := json.Extract(jsonDoc1, path)
		utils.ExpectEqual(expected[i], result.String())
	}
}
