package myexec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/code"
	"github.com/test-instructor/yangfan/hrp/internal/env"
)

var python3Executable string = "python3" // system default python3

func isPython3(python string) bool {
	out, err := Command(python, "--version").Output()
	if err != nil {
		return false
	}
	if strings.HasPrefix(string(out), "Python 3") {
		return true
	}
	return false
}

// EnsurePython3Venv ensures python3 venv with specified packages
// venv should be directory path of target venv
func EnsurePython3Venv(venv string, packages ...string) (python3 string, err error) {
	// priority: specified > $HOME/.hrp/venv
	if venv == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "get user home dir failed")
		}
		venv = filepath.Join(home, ".hrp", "venv")
	}
	python3, err = ensurePython3Venv(venv, packages...)
	if err != nil {
		return "", errors.Wrap(code.InvalidPython3Venv, err.Error())
	}
	python3Executable = python3
	global.GVA_LOG.Info("set python3 executable path", zap.String("Python3Executable", python3Executable))
	return python3, nil
}

func ExecPython3Command(cmdName string, args ...string) error {
	args = append([]string{"-m", cmdName}, args...)
	return RunCommand(python3Executable, args...)
}

func AssertPythonPackage(python3 string, pkgName, pkgVersion string) error {
	out, err := exec.Command(
		python3, "-c", fmt.Sprintf("import %s; print(%s.__version__)", pkgName, pkgName),
	).Output()
	if err != nil {
		return fmt.Errorf("python package %s not found", pkgName)
	}

	// do not check version if pkgVersion is empty
	if pkgVersion == "" {
		global.GVA_LOG.Info("python package is ready", zap.String("name", pkgName))
		return nil
	}

	// check package version equality
	version := strings.TrimSpace(string(out))
	if strings.TrimLeft(version, "v") != strings.TrimLeft(pkgVersion, "v") {
		return fmt.Errorf("python package %s version %s not matched, please upgrade to %s",
			pkgName, version, pkgVersion)
	}

	global.GVA_LOG.Info("python package is ready", zap.String("name", pkgName), zap.String("version", pkgVersion))
	return nil
}

func InstallPythonPackage(python3 string, pkg string) (err error) {
	var pkgName, pkgVersion string
	if strings.Contains(pkg, "==") {
		// funppy==0.5.0
		pkgInfo := strings.Split(pkg, "==")
		pkgName = pkgInfo[0]
		pkgVersion = pkgInfo[1]
	} else {
		// funppy
		pkgName = pkg
	}

	// check if package installed and version matched
	err = AssertPythonPackage(python3, pkgName, pkgVersion)
	if err == nil {
		return nil
	}

	// check if pip available
	err = RunCommand(python3, "-m", "pip", "--version")
	if err != nil {
		global.GVA_LOG.Warn("pip is not available")
		return errors.Wrap(err, "pip is not available")
	}

	global.GVA_LOG.Info("installing python package", zap.String("pkgName", pkgName), zap.String("pkgVersion", pkgVersion))

	// install package
	pypiIndexURL := env.PYPI_INDEX_URL
	if pypiIndexURL == "" {
		pypiIndexURL = "https://pypi.org/simple" // default
	}
	err = RunCommand(python3, "-m", "pip", "install", "--upgrade", pkg,
		"--index-url", pypiIndexURL,
		"--quiet", "--disable-pip-version-check")
	if err != nil {
		return errors.Wrap(err, "pip install package failed")
	}

	return AssertPythonPackage(python3, pkgName, pkgVersion)
}

func RunCommand(cmdName string, args ...string) error {
	cmd := Command(cmdName, args...)
	global.GVA_LOG.Info("exec command", zap.String("cmd", cmd.String()))

	// add cmd dir path to $PATH
	if cmdDir := filepath.Dir(cmdName); cmdDir != "" {
		path := fmt.Sprintf("%s:%s", cmdDir, env.PATH)
		if err := os.Setenv("PATH", path); err != nil {
			global.GVA_LOG.Error("set env $PATH failed", zap.Error(err))
			return err
		}
	}

	// print output with colors
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		global.GVA_LOG.Error("exec command failed", zap.Error(err))
		return err
	}

	return nil
}

func ExecCommandInDir(cmd *exec.Cmd, dir string) error {
	global.GVA_LOG.Info("exec command", zap.String("cmd", cmd.String()), zap.String("dir", dir))
	cmd.Dir = dir

	// print output with colors
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		global.GVA_LOG.Error("exec command failed", zap.Error(err))
		return err
	}

	return nil
}
