package hrp

import (
	"github.com/httprunner/funplugin"
	"github.com/pkg/errors"
	"sync"
)

var cheetahPlugin = make(map[string]*funplugin.IPlugin)
var mutex sync.Mutex

func cheetahInitPlugin(path, venv string, logOn bool) (plugin funplugin.IPlugin, err error) {
	plugins := cheetahPlugin[path]
	if plugins == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if plugins == nil {
			plugin, err = initPlugin(path, venv, logOn)
			if err != nil {
				return nil, errors.Wrap(err, "init plugin failed")
			}
			cheetahPlugin[path] = &plugin
			plugins = &plugin
		}
	}
	return *plugins, nil
}
