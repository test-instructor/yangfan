package interfacecase

import (
	"encoding/json"
	"errors"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"strings"
)

type PyPkgService struct {
}
type PyPkgVersionList struct {
	Version []string `json:"version"`
}

func (p *PyPkgService) PythonEnv() (PyEnvPath string, PipEnvPath string) {
	hostname, _ := os.UserHomeDir()
	PyEnvPath = hostname + "/.hrp/bin/python3"
	PipEnvPath = hostname + "/.hrp/bin/pip3"
	return
}

// todo 后续考虑使用事物，避免出现安装成功，但是数据库未更新的情况

// PyPkgListService 获取Python包列表
func (p *PyPkgService) PyPkgListService(info request.HrpPyPkgRequest) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.HrpPyPkg{})
	var pyPkgLists []interfacecase.HrpPyPkg
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name like ? ", "%"+info.Name+"%")
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&pyPkgLists).Error
	return pyPkgLists, total, err
}

// PyPkgInstallService 安装Python包
func (p *PyPkgService) PyPkgInstallService(pyPkg request.HrpPyPkgRequest, localPkg ...string) (err error) {
	// todo 需要查询一下数据库中是否有改包，避免垃圾数据
	var hrpPyPkg interfacecase.HrpPyPkg
	_, PipEnvPath := p.PythonEnv()
	// 未指定版本号时，安装最新版本
	if pyPkg.Version == "" {
		output, _ := exec.Command(PipEnvPath, "install", pyPkg.Name).Output()
		global.GVA_LOG.Info("安装Python包", zap.String("output", string(output)))
		if strings.Contains(string(output), "Successfully installed") {
			pyPkgInfo, err := p.FindPyPkg(pyPkg.Name)
			if err != nil {
				return err
			}
			pyPkgInfo.IsUninstall = pyPkg.IsUninstall // 入参中是否卸载赋值到pyPkgInfo中
			if p.SelectPyPkg(pyPkg.Name) {
				if err := global.GVA_DB.Model(&hrpPyPkg).Unscoped().Where("name = ?",
					pyPkgInfo.Name).Update("version", pyPkgInfo.Version).Update("deleted_at",
					nil).Error; err != nil {
					return errors.New("入库错误，请验证：" + err.Error())
				}
			} else {
				if err = global.GVA_DB.Create(&pyPkgInfo).Error; err != nil {
					return errors.New("入库错误，请验证：" + err.Error())
				}
			}
			return nil
		} else {
			return errors.New("pip安装出错：" + string(output))
		}
	} else {
		// 指定版本号时，安装指定版本
		output, err := exec.Command(PipEnvPath, "install", pyPkg.Name+"=="+pyPkg.Version).Output()
		if err != nil {
			return err
		}
		if strings.Contains(string(output), "Successfully installed") {
			pyPkgInfo, err := p.FindPyPkg(pyPkg.Name)
			if err != nil {
				return err
			}
			pyPkgInfo.IsUninstall = pyPkg.IsUninstall
			if p.SelectPyPkg(pyPkg.Name) { // 数据库中存在该包
				if err := global.GVA_DB.Model(&hrpPyPkg).Unscoped().Where("name = ?",
					pyPkgInfo.Name).Update("version", pyPkgInfo.Version).Update("deleted_at",
					nil).Error; err != nil {
					return errors.New("入库错误，请验证：" + err.Error())
				}
			} else {
				if err = global.GVA_DB.Create(&pyPkgInfo).Error; err != nil {
					return errors.New("入库错误，请验证：" + err.Error())
				}
			}
			return nil
		} else {
			//global.GVA_LOG.Error("安装Python包失败!", zap.String("pyPkg", string(output)))
			return errors.New("pip安装出错：" + string(output))
		}
	}
}

// UnInstallService 卸载Python包
func (p *PyPkgService) UnInstallService(pkg request.HrpPyPkgRequest) (err error) {
	_, pipEnvPath := p.PythonEnv()
	output, _ := exec.Command(pipEnvPath, "uninstall", pkg.Name, "-y").Output()
	if strings.Contains(string(output), "Successfully uninstalled") {
		err = global.GVA_DB.Where("name = ?", pkg.Name).First(&pkg).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err = global.GVA_DB.Delete(&pkg).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New(string(output))
	}
	return nil
}

// UpdatePyPkgService 更新Python包
func (p *PyPkgService) UpdatePyPkgService(PyPkg request.HrpPyPkgRequest) (err error) {
	_, pipEnvPath := p.PythonEnv()
	var hrpPyPkg interfacecase.HrpPyPkg
	pkgInfo, _ := exec.Command(pipEnvPath, "install", PyPkg.Name, "--upgrade").Output()
	if strings.Contains(string(pkgInfo), "Successfully installed") {
		PyPkgInfo, err := p.FindPyPkg(PyPkg.Name)
		global.GVA_LOG.Info("更新Python包成功!", zap.String("PyPkgInfo.name", PyPkg.Name),
			zap.String("PyPkgInfo.version", PyPkg.Version))
		if err != nil {
			return err
		}
		if err = global.GVA_DB.Model(&hrpPyPkg).Where("name=?", PyPkg.Name).Update("version",
			PyPkgInfo.Version).Error; err != nil {
			global.GVA_LOG.Error(err.Error())
			return err
		}
		return nil
	} else if strings.Contains(string(pkgInfo), "Requirement already satisfied") {
		//global.GVA_LOG.Info("更新Python包失败!", zap.String("pyPkg", string(pkgInfo)))
		if err = global.GVA_DB.Update("version", PyPkg.Version).Error; err != nil {
			return err
		}
		return errors.New("已是最新版本，无需更新")
	} else {
		//global.GVA_LOG.Info("更新Python包失败!", zap.String("pyPkg", string(pkgInfo)))
		return errors.New("更新Python包失败，请检查包名是否正确")
	}

}

// SearchPyPkg 查询Python包--接口
func (p *PyPkgService) SearchPyPkg(pkg request.HrpPyPkgRequest) (pkgInfo request.HrpPyPkgRequest, err error) {
	_, pipEnvPath := p.PythonEnv()
	var pkgList []request.HrpPyPkgRequest
	output, err := exec.Command(pipEnvPath, "list", "--format=json").Output()
	if err != nil {
		return request.HrpPyPkgRequest{}, err
	}
	_ = json.Unmarshal(output, &pkgList)
	for _, pyPkg := range pkgList {
		if strings.ToLower(pyPkg.Name) == strings.ToLower(pkg.Name) {
			scan := global.GVA_DB.First(&pyPkg).Scan(&pyPkg)
			if scan.Error != nil {
				return request.HrpPyPkgRequest{}, scan.Error
			}
			return pyPkg, nil
		}
	}
	return request.HrpPyPkgRequest{}, errors.New("未找到该Python包")
}

// FindPyPkg 查询Python包信息
func (p *PyPkgService) FindPyPkg(name string) (pkgInfo interfacecase.HrpPyPkg, err error) {
	_, PipEnvPath := p.PythonEnv()
	var pkgList []interfacecase.HrpPyPkg
	PyPkgByte, _ := exec.Command(PipEnvPath, "list", "--format=json").Output()
	_ = json.Unmarshal(PyPkgByte, &pkgList)
	for _, pkgInfo = range pkgList {
		if strings.ToLower(pkgInfo.Name) == strings.ToLower(name) {
			//global.GVA_LOG.Info("查询数据库中的python包：", zap.String("入参：", name), zap.String("查询到的信息:", pkgInfo.Name))
			return pkgInfo, nil
		}
	}
	return interfacecase.HrpPyPkg{}, errors.New("未找到该Python包")

}

// SelectPyPkg 查询数据库中是否存在该Python包
func (p *PyPkgService) SelectPyPkg(name string) bool {
	if err := global.GVA_DB.Model(&interfacecase.HrpPyPkg{}).Unscoped().Where("name = ?",
		name).First(&interfacecase.HrpPyPkg{}).Scan(&interfacecase.HrpPyPkg{}).Error; err != nil {
		return false
	}
	return true
}

// PyPkgVersionService 查询Python包版本
func (p *PyPkgService) PyPkgVersionService(pkgInfo request.HrpPyPkgRequest) (err error, list PyPkgVersionList) {
	_, pipEnvPath := p.PythonEnv()
	pkgInfoByte, err := exec.Command(pipEnvPath, "index", "versions", pkgInfo.Name).Output()
	if err != nil {
		return err, list
	}
	if len(pkgInfoByte) == 0 {
		return errors.New("未找到该Python包"), list
	} else {
		split := strings.Split(string(pkgInfoByte), "\n")
		OutVersion := strings.Replace(split[1], "Available versions: ", "", -1)
		pkgVersions := strings.Split(OutVersion, ", ")
		pkgVersions = append(pkgVersions, "")
		copy(pkgVersions[1:], pkgVersions[0:])
		pkgVersions[0] = ""
		// 去除第一个和最后一个空元素
		pkgVersions = pkgVersions[1 : len(pkgVersions)-1]
		list = PyPkgVersionList{
			Version: pkgVersions,
		}
		return nil, list
	}
}
