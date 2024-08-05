package routing

import (
	"testing"
)

func TestICanParseRoutes(t *testing.T) {
	routes := map[string]Route{
		"GET/salut/les/kids": {
			Path:   "/salut/les/kids",
			Method: GET,
		},
		"": {
			Path:   "/",
			Method: ALL,
		},
		"GET/": {
			Path:   "/",
			Method: GET,
		},
	}
	for str, goal := range routes {
		trial := NewFromMethodPathStr(str)
		if trial.Method != goal.Method || trial.Path != goal.Path {
			t.Fatalf("Method (expected: %s, got: %s), Path (expected: %s, got: %s)", goal.Path, trial.Path, goal.Method, trial.Method)
		}
	}
}
