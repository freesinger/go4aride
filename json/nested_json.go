package json

import (
	"strings"
)

func Trim(json string) string {
	return json[1 : len(json)-1]
}

func Split(path string) []string {
	var keys []string
	if len(path) < 2 {
		return nil
	}
	if path[0] == '$' {
		path = path[1:]
		if path[0] == '[' {
			// TODO: resolve array index
		}
		if path[0] == '.' {
			keys = strings.Split(path, ".")
		}
	}
	return keys
}


