package version

import (
	"testing"
	"os"
)

func TestExitWithVersion(t *testing.T) {
	var exited bool
	exit := func(code int) {
		exited = true
	}
	Set("v1.0.0")
	ExitWithVersion([]string{"version_test", "-v"}, exit)
	if !exited {
		t.Errorf("expected exit to have been triggered")
	}
	if v := Get(); v != "v1.0.0" {
		t.Errorf("expected version to be 'v1.0.0', got %v", v)
	}

	exited = false
	ExitWithVersion([]string{"version_test", "--version"}, exit)
	if !exited {
		t.Errorf("expected exit to have been triggered")
	}

	exited = false
	ExitWithVersion([]string{"version_test"}, exit)
	if exited {
		t.Errorf("expected exit to have been skipped")
	}
}

func ExampleExitWithVersion() {
	ExitWithVersion(os.Args, os.Exit)
}
