package server

import "testing"

func Test_Routes_routeNames(t *testing.T) {
	counter := make(map[string]int)

	for _, route := range Routes() {
		name := route.Name

		counter[name]++
		if counter[name] > 1 {
			t.Fatalf("duplicate route name: %q", name)
		}
	}
}

func Test_Routes_routeSignaures(t *testing.T) {
	counter := make(map[string]int)

	for _, route := range Routes() {
		signature := route.Path

		counter[signature]++
		if counter[signature] > 1 {
			t.Fatalf("duplicate route signature: %q", signature)
		}
	}
}
