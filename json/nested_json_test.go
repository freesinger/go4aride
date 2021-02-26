package json

import "testing"

const json = `{
  "bar": {
    "id": 99,
    "mybar": "my mybar"
  },
  "foo": {
    "myfoo": [
      605,
      203,
      910
    ]
  }
}`

func TestTrim(t *testing.T) {
	t.Logf("%v\n", Trim(json))
}