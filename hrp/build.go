package hrp

import (
	"bufio"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/httprunner/funplugin/shared"
	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/hrp/internal/code"
	"github.com/test-instructor/yangfan/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/hrp/internal/version"
)

//go:embed internal/scaffold/templates/plugin/debugtalkPythonTemplate
var pyTemplate string

//go:embed internal/scaffold/templates/plugin/debugtalkGoTemplate
var goTemplate string

// regex for finding all function names
type regexFunctions struct {
	*regexp.Regexp
}

var (
	regexPyFunctionName = regexFunctions{regexp.MustCompile(`(?m)^def ([a-zA-Z_]\w*)\(.*\)`)}
	regexGoFunctionName = regexFunctions{regexp.MustCompile(`(?m)^func ([a-zA-Z_]\w*)\(.*\)`)}
)

func (r *regexFunctions) findAllFunctionNames(content string) ([]string, error) {
	var functionNames []string
	// find all function names
	functionNameSlice := r.FindAllStringSubmatch(content, -1)
	for _, elem := range functionNameSlice {
		name := strings.Trim(elem[1], " ")
		functionNames = append(functionNames, name)
	}

	var filteredFunctionNames []string
	if r == &regexPyFunctionName {
		// filter private functions
		for _, name := range functionNames {
			if strings.HasPrefix(name, "__") {
				continue
			}
			filteredFunctionNames = append(filteredFunctionNames, name)
		}
	} else if r == &regexGoFunctionName {
		// filter main and init function
		for _, name := range functionNames {
			if name == "main" {
				global.GVA_LOG.Warn("plugin debugtalk.go should not define main() function !!!")
				return nil, errors.New("debugtalk.go should not contain main() function")
			}
			if name == "init" {
				continue
			}
			filteredFunctionNames = append(filteredFunctionNames, name)
		}
	}

	global.GVA_LOG.Info("find all function names", zap.Strings("functionNames", filteredFunctionNames))
	return filteredFunctionNames, nil
}

type pluginTemplate struct {
	path          string   // file path
	Version       string   // hrp version
	FunctionNames []string // function names
}

func (pt *pluginTemplate) generate(tmpl, output string) error {
	file, err := os.Create(output)
	if err != nil {
		global.GVA_LOG.Error("open output file failed", zap.Error(err))
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	err = template.Must(template.New("debugtalk").Parse(tmpl)).Execute(writer, pt)
	if err != nil {
		global.GVA_LOG.Error("execute template parsing failed", zap.Error(err))
		return err
	}

	err = writer.Flush()
	if err == nil {
		global.GVA_LOG.Info("generate debugtalk success", zap.String("output", output))
	} else {
		global.GVA_LOG.Info("generate debugtalk failed", zap.Error(err))
	}
	return err
}

func (pt *pluginTemplate) generatePy(output string) error {
	// specify output file path
	if output == "" {
		dir, _ := os.Getwd()
		output = filepath.Join(dir, PluginPySourceGenFile)
	} else if builtin.IsFolderPathExists(output) {
		output = filepath.Join(output, PluginPySourceGenFile)
	}

	// generate .debugtalk_gen.py
	err := pt.generate(pyTemplate, output)
	if err != nil {
		return err
	}

	global.GVA_LOG.Info("build python plugin successfully", zap.String("output", output), zap.String("plugin", pt.path))
	return nil
}

func (pt *pluginTemplate) generateGo(output string) error {
	pluginDir := filepath.Dir(pt.path)
	err := pt.generate(goTemplate, filepath.Join(pluginDir, PluginGoSourceGenFile))
	if err != nil {
		return errors.Wrap(err, "generate hashicorp plugin failed")
	}

	// check go sdk in tempDir
	if err := myexec.RunCommand("go", "version"); err != nil {
		return errors.Wrap(err, "go sdk not installed")
	}

	if !builtin.IsFilePathExists(filepath.Join(pluginDir, "go.mod")) {
		// create go mod
		if err := myexec.ExecCommandInDir(myexec.Command("go", "mod", "init", "main"), pluginDir); err != nil {
			return err
		}

		// download plugin dependency
		// funplugin version should be locked
		funplugin := fmt.Sprintf("github.com/httprunner/funplugin@%s", shared.Version)
		if err := myexec.ExecCommandInDir(myexec.Command("go", "get", funplugin), pluginDir); err != nil {
			return errors.Wrap(err, "go get funplugin failed")
		}
	}

	// add missing and remove unused modules
	if err := myexec.ExecCommandInDir(myexec.Command("go", "mod", "tidy"), pluginDir); err != nil {
		return errors.Wrap(err, "go mod tidy failed")
	}

	// specify output file path
	if output == "" {
		dir, _ := os.Getwd()
		output = filepath.Join(dir, PluginHashicorpGoBuiltFile)
	} else if builtin.IsFolderPathExists(output) {
		output = filepath.Join(output, PluginHashicorpGoBuiltFile)
	}
	outputPath, _ := filepath.Abs(output)

	// build go plugin to debugtalk.bin
	cmd := myexec.Command("go", "build", "-o", outputPath, PluginGoSourceGenFile, filepath.Base(pt.path))
	if err := myexec.ExecCommandInDir(cmd, pluginDir); err != nil {
		return errors.Wrap(err, "go build plugin failed")
	}
	global.GVA_LOG.Info("build go plugin successfully", zap.String("output", outputPath), zap.String("plugin", pt.path))
	return nil
}

// buildGo builds debugtalk.go to debugtalk.bin
func buildGo(path string, output string) error {
	global.GVA_LOG.Info("start to build go plugin", zap.String("path", path), zap.String("output", output))
	content, err := os.ReadFile(path)
	if err != nil {
		global.GVA_LOG.Error("failed to read file", zap.Error(err))
		return errors.Wrap(code.LoadFileError, err.Error())
	}
	functionNames, err := regexGoFunctionName.findAllFunctionNames(string(content))
	if err != nil {
		return errors.Wrap(code.InvalidPluginFile, err.Error())
	}

	templateContent := &pluginTemplate{
		path:          path,
		Version:       version.VERSION,
		FunctionNames: functionNames,
	}
	err = templateContent.generateGo(output)
	if err != nil {
		return errors.Wrap(code.BuildGoPluginFailed, err.Error())
	}
	return nil
}

// buildPy completes funppy information in debugtalk.py
func buildPy(path string, output string) error {
	global.GVA_LOG.Info("start to build python plugin", zap.String("path", path), zap.String("output", output))
	// check the syntax of debugtalk.py
	err := myexec.ExecPython3Command("py_compile", path)
	if err != nil {
		return errors.Wrap(code.InvalidPluginFile,
			fmt.Sprintf("python plugin syntax invalid: %s", err.Error()))
	}

	content, err := os.ReadFile(path)
	if err != nil {
		global.GVA_LOG.Error("failed to read file", zap.Error(err))
		return errors.Wrap(code.LoadFileError, err.Error())
	}
	functionNames, err := regexPyFunctionName.findAllFunctionNames(string(content))
	if err != nil {
		return errors.Wrap(code.InvalidPluginFile, err.Error())
	}

	templateContent := &pluginTemplate{
		path:          path,
		Version:       version.VERSION,
		FunctionNames: functionNames,
	}
	err = templateContent.generatePy(output)
	if err != nil {
		return errors.Wrap(code.BuildPyPluginFailed, err.Error())
	}
	return nil
}

func BuildPlugin(path string, output string) (err error) {
	ext := filepath.Ext(path)
	switch ext {
	case ".py":
		err = buildPy(path, output)
	case ".go":
		err = buildGo(path, output)
	default:
		return errors.Wrap(code.UnsupportedFileExtension,
			"type error, expected .py or .go")
	}
	if err != nil {
		global.GVA_LOG.Error("build plugin failed", zap.Error(err))
		return err
	}
	return nil
}
