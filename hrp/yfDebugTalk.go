package hrp

import (
	"errors"
	"fmt"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
	"os"
	"strconv"
	"sync"
	"time"
)

type debugTalkOperation struct {
	ID        uint
	ProjectID uint
	FilePath  string
	filePath  string
}

func (d *debugTalkOperation) SetFilePath() {
	work, _ := NewWorker(10)
	d.filePath = strconv.Itoa(int(d.ProjectID)) + "_" + strconv.FormatInt(work.NextId(), 10)
}

func (d *debugTalkOperation) CreateDebugTalk(path string) string {
	d.SetFilePath()
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
	d.FilePath = examplesPath + d.filePath
	_, err = os.Stat(d.FilePath)

	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		err := os.Mkdir(d.FilePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
	return d.FilePath
}

func (d *debugTalkOperation) CreateDebugTalkWork() {
	//d.SetFilePath()
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
	_, err = os.Stat(d.FilePath)

	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		err := os.Mkdir(d.FilePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}

func (d *debugTalkOperation) CreateDebugTalkMaster() {
	d.SetFilePath()
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
	d.FilePath = examplesPath + d.filePath
	_, err = os.Stat(d.FilePath)

	//IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		err := os.Mkdir(d.FilePath, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}

func (d *debugTalkOperation) unLock() {
	global.DebugTalkFileLock.RLock()
	global.DebugTalkLock[d.FilePath].Unlock()
	global.DebugTalkFileLock.RUnlock()
}

func (d *debugTalkOperation) getDebugTalkFile() (debugTalkByte []byte, err error) {
	var debugTalkFirst interfacecase.ApiDebugTalk
	db := global.GVA_DB.
		Model(&interfacecase.ApiDebugTalk{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", d.ProjectID)
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
	//d.CreateDebugTalk("")
	global.DebugTalkFileLock.Lock()
	fmt.Println("RunDebugTalkFile:", d.FilePath)
	global.GVA_LOG.Debug("RunDebugTalkFile:" + d.FilePath)
	if global.DebugTalkLock[d.FilePath] == nil {
		global.GVA_LOG.Debug("RunDebugTalkFile:创建锁")
		global.DebugTalkLock[d.FilePath] = &sync.Mutex{}
	}
	global.DebugTalkLock[d.FilePath].Lock()
	global.DebugTalkFileLock.Unlock()

	debugTalkByte, err := d.getDebugTalkFile()
	if err != nil {
		global.GVA_LOG.Error("转换错误", zap.Error(err))
	}
	BuildHashicorpPyPlugin(debugTalkByte, d.FilePath)
}

func (d *debugTalkOperation) StopDebugTalkFile() {
	fmt.Println("StopDebugTalkFile：", d.FilePath)
	fmt.Println("StopDebugTalkFileLock：", global.DebugTalkLock[d.FilePath])
	global.DebugTalkFileLock.RLock()
	global.DebugTalkLock[d.FilePath].Unlock()
	global.DebugTalkFileLock.RUnlock()
	RemoveHashicorpPyPlugin(d.FilePath)
}

const (
	workerBits  uint8 = 10                      //机器码位数
	numberBits  uint8 = 12                      //序列号位数
	workerMax   int64 = -1 ^ (-1 << workerBits) //机器码最大值（即1023）
	numberMax   int64 = -1 ^ (-1 << numberBits) //序列号最大值（即4095）
	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
	workerShift uint8 = numberBits              //机器码偏移量
	epoch       int64 = 1656856144640           //起始常量时间戳（毫秒）,此处选取的时间是2022-07-03 21:49:04
)

type Worker struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("WorkerId超过了限制！")
	}
	return &Worker{
		timeStamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	//当前时间的毫秒时间戳
	now := time.Now().UnixNano() / 1e6
	//如果时间戳与当前时间相同，则增加序列号
	if w.timeStamp == now {
		w.number++
		//如果序列号超过了最大值，则更新时间戳
		if w.number > numberMax {
			for now <= w.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else { //如果时间戳与当前时间不同，则直接更新时间戳
		w.number = 0
		w.timeStamp = now
	}
	//ID由时间戳、机器编码、序列号组成
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}
