package base

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintVersion(t *testing.T) {

	var output bytes.Buffer
	PrintVersion(&output)

	expected := fmt.Sprintf("Version: %s - Commit: %s ", FlunkyVersion, FlunkyCommit)
	assert.Equal(t, expected, output.String())
}
