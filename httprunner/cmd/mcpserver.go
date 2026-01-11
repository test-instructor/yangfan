package cmd

import (
	"github.com/spf13/cobra"
	"github.com/test-instructor/yangfan/httprunner/uixt"
)

var CmdMCPServer = &cobra.Command{
	Use:   "mcp-server",
	Short: "Start MCP server for UI automation",
	Long:  `Start MCP server for UI automation, expose device driver via MCP protocol`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mcpServer := uixt.NewMCPServer()
		return mcpServer.Start()
	},
}
