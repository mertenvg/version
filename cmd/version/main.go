package main

import (
	"os"
	"fmt"

	"github.com/whitecypher/version"
)

var AppVersion string

func main() {
	version.Set(AppVersion)
	version.ExitWithVersion(os.Args, os.Exit)

	fmt.Println("Usage: version -v")
}
