package pkg

import (
	"encoding/json"
	"errors"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
	"time"
)

func PyPkgInstallServiceV2(pyPkg request.HrpPyPkgRequest) (err error) {
	global.GVA_LOG.Debug("安装Python包", zap.Any("pyPkg", pyPkg))
	var hrpPyPkg interfacecase.HrpPyPkg
	// 设置安装参数
	installArgs := []string{"install", pyPkg.Name}
	if pyPkg.Version != "" {
		installArgs[1] = pyPkg.Name + "==" + pyPkg.Version
	}
	hostname, _ := os.UserHomeDir()
	PipEnvPath := hostname + "/.hrp/venv/bin/pip3"
	var getPip = "get-pip.py"
	PyEnvPath := hostname + "/.hrp/venv/bin/python3"
	outputUpdatePip, _ := exec.Command(PyEnvPath, getPip).Output()
	if !strings.Contains(string(outputUpdatePip), "Successfully installed") && !strings.Contains(string(outputUpdatePip), "Successfully uninstalled") {
		global.GVA_LOG.Warn("更新pip失败", zap.String("output", string(outputUpdatePip)))
	}
	global.GVA_LOG.Debug("更新pip信息", zap.String("output", string(outputUpdatePip)))
	output, _ := exec.Command(PipEnvPath, installArgs...).Output()
	global.GVA_LOG.Debug("安装Python包", zap.String("output", string(output)))
	if !strings.Contains(string(output), "Successfully installed") && !strings.Contains(string(output), "Successfully uninstalled") {
		return errors.New("pip安装出错：" + string(output))
	}
	// 更新数据库
	global.GVA_DB.Model(interfacecase.HrpPyPkg{}).First(&hrpPyPkg, "name = ?", pyPkg.Name)
	if hrpPyPkg.ID == 0 {
		hrpPyPkg = interfacecase.HrpPyPkg{}
	}
	hrpPyPkg.Name = pyPkg.Name
	pyPkgInfo, _ := findPyPkgV2(pyPkg.Name, PipEnvPath)
	if pyPkgInfo != nil {
		hrpPyPkg.Version = pyPkgInfo.Version
	} else {
		now := time.Now()
		hrpPyPkg.DeletedAt.Time = now
		hrpPyPkg.DeletedAt.Valid = true
	}
	errSave := global.GVA_DB.Save(&hrpPyPkg).Error
	if errSave != nil {
		return errors.New("入库错误，请验证：" + errSave.Error())
	}
	return nil
}

func findPyPkgV2(name string, PipEnvPath string) (pkgInfo *interfacecase.HrpPyPkg, err error) {
	var pkgList []interfacecase.HrpPyPkg
	PyPkgByte, _ := exec.Command(PipEnvPath, "list", "--format=json").Output()
	_ = json.Unmarshal(PyPkgByte, &pkgList)
	for _, pkg := range pkgList {
		if strings.ToLower(pkg.Name) == strings.ToLower(name) {
			//global.GVA_LOG.Info("查询数据库中的python包：", zap.String("入参：", name), zap.String("查询到的信息:", pkgInfo.Name))
			return &pkg, nil
		}
	}
	return &interfacecase.HrpPyPkg{}, errors.New("未找到该Python包")

}
