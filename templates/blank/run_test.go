package function

import "testing"

// TestRun ensures that the Run function completes without error.
func TestRun(t *testing.T) {
	if err := Run(); err != nil {
		t.Fatal(err)
	}
}
