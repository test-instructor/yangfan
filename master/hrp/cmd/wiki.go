package cmd

import (
	"github.com/spf13/cobra"

	"github.com/test-instructor/yangfan/master/hrp/internal/wiki"
)

var wikiCmd = &cobra.Command{
	Use:     "wiki",
	Aliases: []string{"info", "docs", "doc"},
	Short:   "visit https://httprunner.com",
	PreRun: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return wiki.OpenWiki()
	},
}

func init() {
	rootCmd.AddCommand(wikiCmd)
}
