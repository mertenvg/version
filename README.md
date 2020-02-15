# version
A simple package to deal with application versioning and reporting

# usage
```go
package main

import (
	"os"
	
	"github.com/mertenvg/version"
)

// AppVersion would typically be set at build time using build flags. See example below.
var AppVersion = "v1.0.0"

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