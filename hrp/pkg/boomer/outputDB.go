package boomer

import (
	"encoding/json"
	"github.com/montanaflynn/stats"
	"github.com/rs/zerolog/log"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"strconv"
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
		log.Error().Err(err).Msg("failed to convert data")
		return
	}

	outputStr, _ := json.Marshal(output)
	err = json.Unmarshal(outputStr, &reportDetail)
	pct := map[string]float64{}
	pctList := []float64{75, 80, 85, 90, 95, 99}

	if output.TotalStats != nil {
		var responseTimeData []float64
		for kr, v := range output.TotalStats.ResponseTimes {
			for i := 0; i < int(v); i++ {
				responseTimeData = append(responseTimeData, float64(kr))
			}
		}
		for _, v := range pctList {
			percentile, _ := stats.Percentile(responseTimeData, v)
			pct[strconv.Itoa(int(v))] = percentile
		}
		reportDetail.TotalStats["pct"] = pct
		reportDetail.TotalStats["response_times"] = nil
	}

	if reportDetail.PerformanceReportTotalStats != nil {
		for k, _ := range reportDetail.PerformanceReportTotalStats {
			reportDetail.PerformanceReportTotalStats[k].CurrentRps = output.Stats[k].currentRps
			reportDetail.PerformanceReportTotalStats[k].CurrentFailPerSec = output.Stats[k].currentFailPerSec

			if output.Stats[k].Method != "testcase" && output.Stats[k].Method != "transaction" {
				var responseTimeData []float64
				for kr, v := range output.Stats[k].ResponseTimes {
					for i := 0; i < int(v); i++ {
						responseTimeData = append(responseTimeData, float64(kr))
					}
				}
				output.Stats[k].ResponseTimes = nil
				for _, v := range pctList {
					percentile, _ := stats.Percentile(responseTimeData, v)
					pct[strconv.Itoa(int(v))] = percentile
				}
				reportDetail.PerformanceReportTotalStats[k].ResponseTimer = datatypes.JSONMap{}
				reportDetail.PerformanceReportTotalStats[k].ResponseTimer["pct"] = pct
			}

		}
	}
	if err != nil {
		return
	}
	if o.pTask.State != reportDetail.State {
		o.pTask.State = reportDetail.State
		global.GVA_DB.Save(&o.pTask)
		if reportDetail.State > o.PReport.State {
			o.PReport.State = reportDetail.State
		}
	}

	o.PReport.TotalRPS = output.TotalRPS
	o.PReport.UserCount = output.UserCount
	o.PReport.TotalAvgResponseTime = output.TotalAvgResponseTime
	o.PReport.TotalMinResponseTime = output.TotalMinResponseTime
	o.PReport.TotalMaxResponseTime = output.TotalMaxResponseTime
	o.PReport.TotalFailRatio = output.TotalFailRatio
	o.PReport.TotalFailPerSec = output.TotalFailPerSec
	errReport := global.GVA_DB.Save(&o.PReport).Error
	if errReport != nil {
		global.GVA_LOG.Error("保存测试报告出错", zap.Error(errReport))
		reportString, _ := json.Marshal(o.PReport)
		global.GVA_LOG.Error(string(reportString))
	}
	reportDetail.PerformanceReportID = o.ID
	errReportDetail := global.GVA_DB.Save(&reportDetail).Error
	if errReportDetail != nil {
		global.GVA_LOG.Error("创建性能测试报告详情出错", zap.Error(errReportDetail))
		reportDetailString, _ := json.Marshal(reportDetail)
		global.GVA_LOG.Error(string(reportDetailString))
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
