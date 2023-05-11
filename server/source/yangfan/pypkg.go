package yangfan

import (
	"fmt"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"github.com/test-instructor/yangfan/server/service"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	pyPkgService      = service.ServiceGroupApp.InterfacecaseServiceGroup
	yangfanConfigPath = "temp.yangfan.yaml"
)

type Config struct {
	Pkg *Pkg `json:"pkg" yaml:"pkg"`
}

type PkgInstalledType int

var (
	PkgInstalledTypeFalse PkgInstalledType = 0
	PkgInstalledTypeTrue  PkgInstalledType = 1
)

var _ = []PkgInstalledType{PkgInstalledTypeFalse, PkgInstalledTypeTrue}

type Pkg struct {
	Installed PkgInstalledType `json:"installed" yaml:"installed"`
}

func PyPkg() {
	var config Config
	bytes, err := os.ReadFile(yangfanConfigPath)
	if err != nil {
		global.GVA_LOG.Warn(fmt.Sprintln("扬帆配置文件不存在:", yangfanConfigPath))
		config.Pkg = new(Pkg)
		config.Pkg.Installed = PkgInstalledTypeFalse
		configOutput, err := yaml.Marshal(config)
		if err != nil {
			global.GVA_LOG.Error("config 转换失败：", zap.Error(err))
			return
		}
		if err := os.WriteFile(yangfanConfigPath, configOutput, 0644); err != nil {
			global.GVA_LOG.Error("配置文件写入失败：", zap.Error(err))
		}
	} else {
		err = yaml.Unmarshal(bytes, &config)
		if err != nil {
			global.GVA_LOG.Warn("读取配置文件失败:", zap.Error(err))
		}
		if config.Pkg != nil && config.Pkg.Installed == PkgInstalledTypeTrue {
			global.GVA_LOG.Info("初始化时插件已经安装")
			return
		}
	}

	var packages []interfacecase.HrpPyPkg

	db := global.GVA_DB.Model(interfacecase.HrpPyPkg{})
	err = db.Find(&packages).Error
	if err != nil {
		global.GVA_LOG.Error("获取 python 第三方库失败", zap.Error(err))
		return
	}
	for _, pyPkg := range packages {
		global.GVA_LOG.Debug("安装python插件")
		if err := pyPkgService.PyPkgInstallService(request.HrpPyPkgRequest{HrpPyPkg: pyPkg}); err != nil {
			global.GVA_LOG.Error("安装 python 第三方库失败", zap.Any("pyPkg", pyPkg), zap.Error(err))
		}
	}

	config.Pkg.Installed = PkgInstalledTypeTrue
	configOutput, err := yaml.Marshal(config)
	if err != nil {
		global.GVA_LOG.Error("config 转换失败：", zap.Error(err))
		return
	}
	if err := os.WriteFile(yangfanConfigPath, configOutput, 0644); err != nil {
		global.GVA_LOG.Error("配置文件写入失败：", zap.Error(err))
	}
}
