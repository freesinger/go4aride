package json

import (
	"log"
	"testing"
)

func ExpectEqual(exp interface{}, real interface{}) {
	switch exp.(type) {
	case int:
		if exp.(int) != real.(int) {
			log.Fatalf("expected %v, but got %v", exp, real)
		}
	case string:
		if exp.(string) != real.(string) {
			log.Fatalf("expected %v, but got %v", exp, real)
		}
	}
	log.Printf("expected %v, got %v", exp, real)
}

//func TestTrim(t *testing.T) {
//	t.Logf("%v\n", Trim(json))
//	t.Logf("%v\n", 10^10)
//	ExpectEqual(10 ^ 10, 0)
//}

func TestGet(t *testing.T) {
	json := `{"bar": {"id": 99, "mybar": "my mybar" }, "foo": ` +
		`{"myfoo": [605,203,910]}}`
	paths := []string{"foo.myfoo", "bar.id", "bar.xyz", "bar.mybar"}
	expected := []string{"[605,203,910]", "99", "", "my mybar"}
	results := GetMany(json, paths...)
	ExpectEqual(len(results), len(paths))
	// log.Printf("%v\n", gjson.GetMany(json, paths...))
	for i, _ := range paths {
		ExpectEqual(expected[i], results[i].String())
	}
}

func TestString(t *testing.T) {

}
