package base

import (
	"fmt"
	"io"
)

var FlunkyVersion = "development"
var FlunkyCommit = "dev"

func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "Version: %s - Commit: %s ", FlunkyVersion, FlunkyCommit)
}
