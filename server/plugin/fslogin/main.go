package fslogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cp "github.com/otiai10/copy"
	global2 "github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/global"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/model"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/router"
	"os"
	"path/filepath"
	"strings"
)

type FsLoginPlugin struct {
}

func CreateFsLoginPlug(engine *gin.Engine) *FsLoginPlugin {

	global.GlobalConfig = global2.GVA_CONFIG.FS

	global2.GVA_DB.AutoMigrate(model.FsUserInfo{})
	_, err := os.Stat("./resource/fsView")
	if err != nil {
		if os.IsNotExist(err) {
			cp.Copy("./plugin/fslogin/resource", "./resource")
			fmt.Println("飞书登录oAuth2.0插件资源文件夹初始化完成，资源目录为：./resource/fsView")
		}
	} else {
		fmt.Println("resource/fsView已存在，如已安装过飞书登录插件，请忽略此提示")
	}

	var files []string
	filepath.Walk("./resource", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	engine.LoadHTMLFiles(files...)
	return &FsLoginPlugin{}
}

func (*FsLoginPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitFsLoginRouter(group)
}

func (*FsLoginPlugin) RouterPath() string {
	return "fsLogin"
}
