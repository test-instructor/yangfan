package runTestCase

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"go.uber.org/zap"
)

// Adapted from yangfan/run/runTestCase/debugTalk.go

// Mocking global locks if not present in yangfan global
var (
	DebugTalkFileLock sync.RWMutex
	DebugTalkLock     = make(map[string]*sync.Mutex)
)

type debugTalkOperation struct {
	ID        uint
	ProjectID uint
	FilePath  string
}

func (d *debugTalkOperation) CreateDebugTalk(path string) string {
	examplesPath := "examples/"
	_, err := os.Stat(examplesPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(examplesPath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	filePath := examplesPath + path
	_, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		err := os.Mkdir(filePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	return filePath
}

func (d *debugTalkOperation) unLock() {
	DebugTalkFileLock.RLock()
	DebugTalkLock[d.FilePath].Unlock()
	DebugTalkFileLock.RUnlock()
}

func (d *debugTalkOperation) getDebugTalkFile(projectID uint) (debugTalkByte []byte, err error) {
	var debugTalkFirst platform.PythonCode
	db := global.GVA_DB.Model(&platform.PythonCode{}).
		Where("project_id = ?", projectID)
	// Type = 1 for automation
	db.Where("type = ?", 1).Order("id desc")
	err = db.First(&debugTalkFirst).Error
	if err != nil {
		// Fallback or return error
		return nil, err
	}
	return []byte(debugTalkFirst.Code), err
}

func (d *debugTalkOperation) RunDebugTalkFile() error {
	debugTalkByte, err := d.getDebugTalkFile(d.ProjectID)
	if err != nil {
		global.GVA_LOG.Error("getDebugTalkFile failed", zap.Error(err))
		return err
	}

	d.FilePath = d.CreateDebugTalk(fmt.Sprintf("TdebugTalk_%d_%d/", d.ID, rand.Int31n(99999999)))
	DebugTalkFileLock.Lock()
	fmt.Println("RunDebugTalkFile:", d.FilePath)
	global.GVA_LOG.Debug("RunDebugTalkFile:" + d.FilePath)
	if DebugTalkLock[d.FilePath] == nil {
		global.GVA_LOG.Debug("RunDebugTalkFile:创建锁")
		DebugTalkLock[d.FilePath] = &sync.Mutex{}
	}
	DebugTalkLock[d.FilePath].Lock()
	DebugTalkFileLock.Unlock()

	hrp.BuildHashicorpPyPlugin(debugTalkByte, d.FilePath)
	return nil
}

func (d *debugTalkOperation) StopDebugTalkFile() {
	fmt.Println("StopDebugTalkFile：", d.FilePath)
	fmt.Println("StopDebugTalkFileLock：", DebugTalkLock[d.FilePath])
	DebugTalkFileLock.RLock()
	DebugTalkLock[d.FilePath].Unlock()
	DebugTalkFileLock.RUnlock()
	hrp.RemoveHashicorpPyPlugin(d.FilePath)
}
