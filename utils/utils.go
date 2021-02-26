package utils

import "log"

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
