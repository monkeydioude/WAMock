package routing

import "testing"

func TestICanFindMethodFromAString(t *testing.T) {
	tests := map[string]Method{
		"ALL":    ALL,
		"sasas":  NONE,
		"GET":    GET,
		"":       NONE,
		"PUT":    PUT,
		"DELETE": DELETE,
		"PATCH":  PATCH,
	}
	for str, goal := range tests {
		trial, _ := SeekMethod(str)
		if !goal.MatchString(trial.String()) {
			t.Fatalf("%s != %s", trial, goal)
		}
	}
}
