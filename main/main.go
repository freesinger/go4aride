package main

import (
	"github.com/freesinger/go4aride/json"
	"github.com/freesinger/go4aride/utils"
)

func main() {
	jsonDoc := `{"bar": {"id": 99, "mybar": "my mybar" }, "foo": ` +
		`{"myfoo": [605,203,910]}}`
	paths := []string{"foo.myfoo", "bar.id", "bar.xyz", "bar.mybar"}
	expected := []string{"[605,203,910]", "99", "", "my mybar"}
	//results := json.GetMany(jsonDoc, paths...)
	//utils.ExpectEqual(len(results), len(paths))
	// log.Printf("%v\n", gjson.GetMany(json, paths...))
	for i, path := range paths {
		result := json.Get(jsonDoc, path)
		utils.ExpectEqual(expected[i], result.String())
	}
}
