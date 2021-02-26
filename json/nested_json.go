package json
//
//import "strings"
//
//type Type int
//
//const (
//	Null Type = iota
//	True
//	False
//	Number
//	String
//	JSON
//)
//
//type Result struct {
//	Type  Type
//	Raw   string
//	Str   string
//	Num   float64
//	Index int
//}
//
//func Trim(json string) string {
//	return json[1 : len(json)-1]
//}
//
//func Extract(json, path string) Result {
//	var keys []string
//	var depth int
//	json = Trim(json)
//	if len(path) > 2 {
//		depth = 1
//		if path[0:1] == "$." {
//			path = path[2:]
//			keys = strings.Split(path, ".")
//		}
//	}
//	for i := 0; i < len(keys); i++ {
//		s := keys[i]
//		var start, end int
//		// TODO: s has array struct
//		for j := 0; j < len(json); j++ {
//			switch json[j] {
//			case '"':
//				j++
//				start = j
//				// 抽象一个取字符串值的函数出来
//				for ; j < len(json); j++ {
//					if json[j] == '"' {
//						end = j
//						if json[start:end] == s && depth == i+1 {
//							if i != len(keys)-1 {
//								json = json[end+1:]
//								break
//							} else {
//								//
//								return Result{}
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//
//	return Result{}
//}
