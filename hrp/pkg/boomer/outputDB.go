package boomer

import (
	"encoding/json"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

type DbOutput struct {
	ID            uint
	pTask         interfacecase.Performance
	PReport       interfacecase.PerformanceReport
	ReportChannel chan map[string]interface{}
	closeChannel  bool
}

// NewDbOutput returns a ConsoleOutput.
func NewDbOutput(pReport interfacecase.PerformanceReport, pTask interfacecase.Performance) *DbOutput {
	return &DbOutput{ID: pReport.ID, pTask: pTask, PReport: pReport}
}

func NewDbOutputWork(reportID uint, taskID uint) *DbOutput {
	var pTask interfacecase.Performance
	var pReport interfacecase.PerformanceReport
	global.GVA_DB.Model(interfacecase.Performance{}).First(&pTask, "id = ?", taskID)
	global.GVA_DB.Model(interfacecase.PerformanceReport{}).First(&pReport, "id = ?", reportID)
	return &DbOutput{ID: reportID, pTask: pTask, PReport: pReport}
}

func (o *DbOutput) OnStart() {
	o.ReportChannel = make(chan map[string]interface{}, 1000)
	go func() {
		for {
			if o.closeChannel == true {
				break
			} else {
				select {
				case data, ok := <-o.ReportChannel:
					o.updateReport(data)
					if !ok {
						break
					}
				}
			}

		}
		close(o.ReportChannel)
	}()
}

// OnStop of DbOutput has nothing to do.
func (o *DbOutput) OnStop() {
}

func (o *DbOutput) updateReport(data map[string]interface{}) {
	var reportDetail interfacecase.PerformanceReportDetail
	output, err := convertData(data)
	if err != nil {
		global.GVA_LOG.Error("failed to convert data", zap.Error(err))
		return
	}

	outputStr, _ := json.Marshal(output)
	err = json.Unmarshal(outputStr, &reportDetail)

	if o.pTask.State != reportDetail.State {
		o.pTask.State = reportDetail.State
		global.GVA_DB.Save(&o.pTask)
		if reportDetail.State > o.PReport.State {
			o.PReport.State = reportDetail.State
		}
	}

	errReport := global.GVA_DB.Save(&o.PReport).Error
	if errReport != nil {
		global.GVA_LOG.Error("保存测试报告出错", zap.Error(errReport))
		reportString, _ := json.Marshal(o.PReport)
		global.GVA_LOG.Error(string(reportString))
	}
	if o.PReport.State == interfacecase.StateStopped {
		o.closeChannel = true
	}
	return
}

// OnEvent will print to the console.
func (o *DbOutput) OnEvent(data map[string]interface{}) {
	if !o.closeChannel {
		o.ReportChannel <- data
	}
}
