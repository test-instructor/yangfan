package initialize

import (
	"fmt"
	"strconv"

	"github.com/robfig/cron/v3"

	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"

	"github.com/test-instructor/yangfan/server/config"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GVA_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}

func TimerTaskCase() {
	var timerTaskCase []interfacecase.ApiTimerTask
	global.GVA_DB.Model(interfacecase.ApiTimerTask{}).
		Where("Status = ?", true).
		Find(&timerTaskCase)
	for _, task := range timerTaskCase {
		id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(task.ID)), task.RunTime, runTestCase.RunTimerTaskBack(task.ID), cron.WithSeconds())
		if err != nil {
			return
		}
		task.EntryID = int(id)
		err = global.GVA_DB.Save(&task).Error
		if err != nil {
			return
		}
	}
}
