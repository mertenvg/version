package main

import (
	"os"
	"fmt"

	"github.com/whitecypher/version"
)

var (
	AppVersion = "v0.0.0"
)

func main() {
	version.Set(AppVersion)
	version.ExitWithVersion(os.Args, os.Exit)

	fmt.Println("Usage: version -v")
}
