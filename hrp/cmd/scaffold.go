package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/scaffold"
)

var scaffoldCmd = &cobra.Command{
	Use:     "startproject $project_name",
	Aliases: []string{"scaffold"},
	Short:   "create a scaffold project",
	Args:    cobra.ExactValidArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ignorePlugin && !genPythonPlugin && !genGoPlugin {
			return errors.New("please specify function plugin type")
		}

		var pluginType scaffold.PluginType
		if empty {
			pluginType = scaffold.Empty
		} else if ignorePlugin {
			pluginType = scaffold.Ignore
		} else if genGoPlugin {
			pluginType = scaffold.Go
		} else {
			pluginType = scaffold.Py // default
		}

		err := scaffold.CreateScaffold(args[0], pluginType, venv, force)
		if err != nil {
			global.GVA_LOG.Error("create scaffold project failed", zap.Error(err))
			return err
		}
		global.GVA_LOG.Info("create scaffold project success", zap.String("projectName", args[0]))
		return nil
	},
}

var (
	empty           bool
	ignorePlugin    bool
	genPythonPlugin bool
	genGoPlugin     bool
	force           bool
)

func init() {
	rootCmd.AddCommand(scaffoldCmd)
	scaffoldCmd.Flags().BoolVarP(&force, "force", "f", false, "force to overwrite existing project")
	scaffoldCmd.Flags().BoolVar(&genPythonPlugin, "py", true, "generate hashicorp python plugin")
	scaffoldCmd.Flags().BoolVar(&genGoPlugin, "go", false, "generate hashicorp go plugin")
	scaffoldCmd.Flags().BoolVar(&ignorePlugin, "ignore-plugin", false, "ignore function plugin")
	scaffoldCmd.Flags().BoolVar(&empty, "empty", false, "generate empty project")
}
