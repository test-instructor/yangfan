package runTestCase

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type debugTalkOperation struct {
	ID        uint
	ProjectID uint
	FilePath  string
}

func (d *debugTalkOperation) CreateDebugTalk(path string) string {
	examplesPath := "examples/"
	_, err := os.Stat(examplesPath)
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		err := os.Mkdir(examplesPath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	filePath := examplesPath + path
	_, err = os.Stat(filePath)

	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
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
	global.DebugTalkFileLock.RLock()
	global.DebugTalkLock[d.FilePath].Unlock()
	global.DebugTalkFileLock.RUnlock()
}

func (d *debugTalkOperation) getDebugTalkFile(projectID uint) (debugTalkByte []byte, err error) {
	var debugTalkFirst interfacecase.ApiDebugTalk
	db := global.GVA_DB.
		Model(&interfacecase.ApiDebugTalk{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", projectID)
	//查询对应的类型
	db.Where("file_type = ?", interfacecase.FileDebugTalk).Order("id desc")
	err = db.First(&debugTalkFirst).Error
	if err != nil {
		defaultDB := global.GVA_DB.Model(&interfacecase.ApiDebugTalk{}).
			Preload("Project").Joins("Project").Where("Project.ID = ?", 1)
		defaultDB.Where("file_type = ?", interfacecase.FileDebugTalk)
		err = defaultDB.First(&debugTalkFirst).Error
	}
	return []byte(debugTalkFirst.Content), err
}

func (d *debugTalkOperation) RunDebugTalkFile() {
	d.FilePath = d.CreateDebugTalk(fmt.Sprintf("TdebugTalk_%d_%d/", d.ID, rand.Int31n(99999999)))
	global.DebugTalkFileLock.Lock()
	fmt.Println("RunDebugTalkFile:", d.FilePath)
	global.GVA_LOG.Debug("RunDebugTalkFile:" + d.FilePath)
	if global.DebugTalkLock[d.FilePath] == nil {
		global.GVA_LOG.Debug("RunDebugTalkFile:创建锁")
		global.DebugTalkLock[d.FilePath] = &sync.Mutex{}
	}
	global.DebugTalkLock[d.FilePath].Lock()
	global.DebugTalkFileLock.Unlock()

	debugTalkByte, _ := d.getDebugTalkFile(d.ProjectID)
	hrp.BuildHashicorpPyPlugin(debugTalkByte, d.FilePath)
}

func (d *debugTalkOperation) StopDebugTalkFile() {
	fmt.Println("StopDebugTalkFile：", d.FilePath)
	fmt.Println("StopDebugTalkFileLock：", global.DebugTalkLock[d.FilePath])
	global.DebugTalkFileLock.RLock()
	global.DebugTalkLock[d.FilePath].Unlock()
	global.DebugTalkFileLock.RUnlock()
	hrp.RemoveHashicorpPyPlugin(d.FilePath)
}
