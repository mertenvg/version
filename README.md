# version
A simple package to deal with application versioning and reporting

# usage
```go
package main

import (
	"os"
	
	"github.com/whitecypher/version"
)

var (
	AppVersion = "v1.0.0"
)

func main() {
    version.Set(AppVersion)
    version.ExitWithVersion(os.Args, os.Exit)
}
```
you can optionally automatically set AppVersion while building using the `ldflags` option;
```bash
go install -ldflags "-w -X main.AppVersion=$(git describe --tags)" ./cmd/version
version -v
```