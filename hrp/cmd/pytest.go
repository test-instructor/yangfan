package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/hrp/internal/pytest"
	"github.com/test-instructor/yangfan/hrp/internal/version"
)

var pytestCmd = &cobra.Command{
	Use:   "pytest $path ...",
	Short: "run API test with pytest",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
	},
	DisableFlagParsing: true, // allow to pass any args to pytest
	RunE: func(cmd *cobra.Command, args []string) error {
		packages := []string{
			fmt.Sprintf("httprunner==%s", version.HttpRunnerMinimumVersion),
		}
		_, err := myexec.EnsurePython3Venv(venv, packages...)
		if err != nil {
			global.GVA_LOG.Error("python3 venv is not ready", zap.Error(err))
			return err
		}
		return pytest.RunPytest(args)
	},
}

func init() {
	rootCmd.AddCommand(pytestCmd)
}
