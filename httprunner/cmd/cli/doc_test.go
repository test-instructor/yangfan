package main

import (
	"testing"

	"github.com/spf13/cobra/doc"
	"github.com/stretchr/testify/assert"
	"github.com/test-instructor/yangfan/httprunner/cmd"
)

// run this test to generate markdown docs for hrp command
func TestGenMarkdownTree(t *testing.T) {
	addAllCommands()
	err := doc.GenMarkdownTree(cmd.RootCmd, "../../docs/cmd")
	assert.Nil(t, err)
}
