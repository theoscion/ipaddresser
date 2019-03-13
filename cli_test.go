package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCLI(test *testing.T) {
	test.Parallel()

	expectedCLI := &CLI{
		Verbose:    true,
		Daemon:     true,
		Hooks:      map[string]int{"http://localhost/": 0},
		AlwaysHook: true,
	}

	receivedFullCLI := NewCLI([]string{"--verbose", "--daemon", "--always-hook", "http://localhost/"})
	receivedShortCLI := NewCLI([]string{"-v", "-d", "-a", "http://localhost/"})

	assert.Equal(test, expectedCLI, receivedFullCLI, "NewCLI() expected to match")
	assert.Equal(test, expectedCLI, receivedShortCLI, "NewCLI() expected to match")
}
