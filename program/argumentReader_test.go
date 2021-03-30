package program

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadProgramArgument(t *testing.T) {
	origArgs := os.Args
	os.Args = append(os.Args, "-value", "val")

	value, err := ReadProgramArgument("value")
	assert.NoError(t, err)
	assert.Equal(t, value, "val")

	os.Args = origArgs
}

func TestReadProgramArgumentWithoutValue(t *testing.T) {
	origArgs := os.Args
	os.Args = append(os.Args, "-value")

	_, err := ReadProgramArgument("value")

	assert.Error(t, err, "value not found")

	os.Args = origArgs
}

func TestReadEmptyProgramArgument(t *testing.T) {
	origArgs := os.Args

	_, err := ReadProgramArgument("value")
	assert.Error(t, err, "argument not found")

	os.Args = origArgs
}

func TestReadProgramFlag(t *testing.T) {
	origArgs := os.Args

	os.Args = append(os.Args, "-flag")

	value := ReadProgramFlag("flag")
	assert.Equal(t, value, true)

	os.Args = origArgs
}

func TestReadEmptyProgramFlag(t *testing.T) {
	origArgs := os.Args

	value := ReadProgramFlag("flag")
	assert.Equal(t, value, false)

	os.Args = origArgs
}
