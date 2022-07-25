package base

import (
	"fmt"
)

var FlunkyVersion = "development"
var FlunkyCommit = "dev"

func PrintVersion() {
	fmt.Printf("Version: %s - Commit: %s ", FlunkyVersion, FlunkyCommit)
}
