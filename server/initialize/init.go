// 假设这是初始化逻辑的一部分

package initialize

import (
	"github.com/test-instructor/yangfan/server/v2/utils"
)

// 初始化全局函数
func SetupHandlers() {
	// 注册系统重载处理函数
	utils.GlobalSystemEvents.RegisterReloadHandler(func() error {
		return Reload()
	})
}
