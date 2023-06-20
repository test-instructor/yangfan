package hrp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/httprunner/funplugin"
	"github.com/httprunner/funplugin/fungo"
	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/code"
	"github.com/test-instructor/yangfan/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/hrp/internal/sdk"
)

const (
	PluginGoBuiltFile          = "debugtalk.so"      // built from go official plugin
	PluginHashicorpGoBuiltFile = "debugtalk.bin"     // built from hashicorp go plugin
	PluginGoSourceFile         = "debugtalk.go"      // golang function plugin source file
	PluginGoSourceGenFile      = "debugtalk_gen.go"  // generated for hashicorp go plugin
	PluginPySourceFile         = "debugtalk.py"      // python function plugin source file
	PluginPySourceGenFile      = ".debugtalk_gen.py" // generated for hashicorp python plugin
)

const projectInfoFile = "proj.json" // used for ensuring root project

var pluginMap sync.Map // used for reusing plugin instance
var pluginMapRW = make(map[string]*funplugin.IPlugin)
var pluginMutex sync.RWMutex

func initPlugin(path, venv string, logOn bool) (plugin funplugin.IPlugin, err error) {
	pluginMutex.RLock()
	plugins := pluginMapRW[path]
	pluginMutex.RUnlock()
	if plugins == nil {
		pluginMutex.Lock()
		defer pluginMutex.Unlock()
		plugins = pluginMapRW[path]
		if plugins == nil {
			plugin, err = initplugin(path, venv, logOn)
			if err != nil {
				return nil, errors.Wrap(err, "init plugin failed")
			}
			pluginMapRW[path] = &plugin
			plugins = &plugin
		}
	}
	return *plugins, nil
}

func initplugin(path, venv string, logOn bool) (plugin funplugin.IPlugin, err error) {
	// plugin file not found
	if path == "" {
		return nil, nil
	}
	pluginPath, err := locatePlugin(path)
	if err != nil {
		global.GVA_LOG.Warn("locate plugin failed", zap.String("path", path), zap.Error(err))
		return nil, nil
	}

	// reuse plugin instance if it already initialized
	if p, ok := pluginMap.Load(pluginPath); ok {
		return p.(funplugin.IPlugin), nil
	}

	pluginOptions := []funplugin.Option{funplugin.WithLogOn(logOn)}

	if strings.HasSuffix(pluginPath, ".py") {
		// register funppy plugin
		genPyPluginPath := filepath.Join(filepath.Dir(pluginPath), PluginPySourceGenFile)
		err = BuildPlugin(pluginPath, genPyPluginPath)
		if err != nil {
			global.GVA_LOG.Error("build plugin failed", zap.String("path", pluginPath), zap.Error(err))
			return nil, err
		}
		pluginPath = genPyPluginPath

		packages := []string{
			fmt.Sprintf("funppy==%s", fungo.Version),
		}
		python3, err := myexec.EnsurePython3Venv(venv, packages...)
		if err != nil {
			global.GVA_LOG.Error("python3 venv is not ready", zap.Any("packages", packages), zap.Error(err))
			return nil, err
		}
		pluginOptions = append(pluginOptions, funplugin.WithPython3(python3))
	}

	// found plugin file
	plugin, err = funplugin.Init(pluginPath, pluginOptions...)
	if err != nil {
		global.GVA_LOG.Error("init plugin failed", zap.String("path", pluginPath), zap.Error(err))
		err = errors.Wrap(code.InitPluginFailed, err.Error())
		return
	}

	// add plugin instance to plugin map
	pluginMap.Store(pluginPath, plugin)

	// report event for initializing plugin
	event := sdk.EventTracking{
		Category: "InitPlugin",
		Action:   fmt.Sprintf("Init %s plugin", plugin.Type()),
		Value:    0, // success
	}
	if err != nil {
		event.Value = 1 // failed
	}
	go sdk.SendEvent(event)

	return
}

func locatePlugin(path string) (pluginPath string, err error) {
	// priority: hashicorp plugin (debugtalk.bin > debugtalk.py) > go plugin (debugtalk.so)

	pluginPath, err = locateFile(path, PluginHashicorpGoBuiltFile)
	if err == nil {
		return
	}

	pluginPath, err = locateFile(path, PluginPySourceFile)
	if err == nil {
		return
	}

	pluginPath, err = locateFile(path, PluginGoBuiltFile)
	if err == nil {
		return
	}

	return "", fmt.Errorf("plugin file not found")
}

// locateFile searches destFile upward recursively until system root dir
func locateFile(startPath string, destFile string) (string, error) {
	stat, err := os.Stat(startPath)
	if os.IsNotExist(err) {
		return "", err
	}

	var startDir string
	if stat.IsDir() {
		startDir = startPath
	} else {
		startDir = filepath.Dir(startPath)
	}
	startDir, _ = filepath.Abs(startDir)

	// convention over configuration
	pluginPath := filepath.Join(startDir, destFile)
	if _, err := os.Stat(pluginPath); err == nil {
		return pluginPath, nil
	}

	// system root dir
	parentDir, _ := filepath.Abs(filepath.Dir(startDir))
	if parentDir == startDir {
		return "", fmt.Errorf("searched to system root dir, plugin file not found")
	}

	return locateFile(parentDir, destFile)
}

func GetProjectRootDirPath(path string) (rootDir string, err error) {
	pluginPath, err := locatePlugin(path)
	if err == nil {
		rootDir = filepath.Dir(pluginPath)
		return
	}
	// fix: no debugtalk file in project but having proj.json created by startpeoject
	projPath, err := locateFile(path, projectInfoFile)
	if err == nil {
		rootDir = filepath.Dir(projPath)
		return
	}

	// failed to locate project root dir
	// maybe project plugin debugtalk.xx and proj.json are not exist
	// use current dir instead
	return os.Getwd()
}
