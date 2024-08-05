package config

import "testing"

func TestICanParseArgs(t *testing.T) {
	trials := []RunConfig{
		RetrieveStartingConf([]string{"zdedededex", "run_config_test.go", "-x", "20"}),
		RetrieveStartingConf([]string{":)", "."}),
	}
	goals := []RunConfig{{
		isDirectory:           false,
		confPath:              "run_config_test.go",
		coroutineRefreshTimer: 20,
	}, {
		isDirectory:           true,
		confPath:              ".",
		coroutineRefreshTimer: 0,
	}}

	for it, trial := range trials {
		if trial.isDirectory != goals[it].isDirectory ||
			trial.confPath != goals[it].confPath ||
			trial.coroutineRefreshTimer != goals[it].coroutineRefreshTimer {
			t.Fatalf("it: %d, expected: %+v, got: %+v\n", it, trial, goals[it])
		}
	}
}
