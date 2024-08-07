package file_system

import "testing"

func TestICanSanitizeFilePath(t *testing.T) {
	trials := []string{"GET:test:1.json"}
	goals := []string{"GET/test/1"}

	for i, trial := range trials {
		if CleanConfigFilename(trial) != goals[i] {
			t.Fail()
		}
	}
}
