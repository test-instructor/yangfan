package hrp

import (
	"github.com/httprunner/funplugin"
	"github.com/pkg/errors"
	"sync"
)

var pluginMapRW = make(map[string]*funplugin.IPlugin)
var pluginMutex sync.RWMutex

func yangfanInitPlugin(path, venv string, logOn bool) (plugin funplugin.IPlugin, err error) {
	pluginMutex.RLock()
	plugins := pluginMapRW[path]
	pluginMutex.RUnlock()
	if plugins == nil {
		pluginMutex.Lock()
		defer pluginMutex.Unlock()
		plugins = pluginMapRW[path]
		if plugins == nil {
			plugin, err = initPlugin(path, venv, logOn)
			if err != nil {
				return nil, errors.Wrap(err, "init plugin failed")
			}
			pluginMapRW[path] = &plugin
			plugins = &plugin
		}
	}
	return *plugins, nil
}
