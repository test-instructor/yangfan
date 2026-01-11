package initialize

import (
	"os"
	"path/filepath"

	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

func InitPython() {
	home, err := os.UserHomeDir()
	if err != nil {
		global.GVA_LOG.Error("get user home dir fail", zap.Error(err))
	}
	venv := filepath.Join(home, ".lc", "venv")
	global.PythonVENV = filepath.Join(venv, "bin", "python3")
	//var pk []platform.PythonPackage
	//var pks []string
	//pks = append(pks, "funppy==0.5.0")
	//global.GVA_DB.Find(&pk)
	//for _, v := range pk {
	//	pks = append(pks, fmt.Sprintf("%s==%s", *v.Name, v.Version))
	//}
	//_, err = myexec.EnsurePython3Venv(venv)
	//if err != nil {
	//	global.GVA_LOG.Error("init python venv fail", zap.Error(err))
	//	err = myexec.RunCommand("python3", "-m", "venv", venv)
	//	if err != nil {
	//		global.GVA_LOG.Error("init python venv fail", zap.Error(err))
	//	}
	//	_, err = myexec.EnsurePython3Venv(venv)
	//	if err != nil {
	//		global.GVA_LOG.Error("init python venv fail", zap.Error(err))
	//	}
	//}
	//myexec.InstallPip(global.PythonVENV)
	//for _, pkg := range pks {
	//	myexec.InstallPythonPackage(global.PythonVENV, pkg)
	//}
}
