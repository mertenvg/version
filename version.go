// version attempts to reduce version option handling logic between application while remaining as loosely coupled as
// possible from your architecture.
package version

import (
	"fmt"
)

var (
	v string
)

type ExitFunc func (code int)

// ExitWithVersion scans the given args for an occurrence of -v or --version and if found, prints the version and exits.
// This function assumes the arguments come from os.Args where the first argument is the name of the binary being executed.
func ExitWithVersion(args []string, exit ExitFunc) {
	for _, o := range args {
		if o == "-v" || o == "--version" {
			fmt.Printf("%s %s\n", args[0], v)
			exit(0)
		}
	}
}

// Set the version
func Set(version string) {
	v = version
}

// Get the version
func Get() string {
	return v
}

// String is an alias to version.Get().
func String() string {
	return Get()
}