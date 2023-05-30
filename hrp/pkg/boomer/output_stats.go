package boomer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

var (
	instanceMaster = "master"
	instanceWorker = "worker"
)

var (
	masterGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "performance_report_master",
		Help: "Performance report for master",
	}, []string{instanceMaster})
	workGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "performance_report_work",
		Help: "Performance report for work",
	}, []string{instanceWorker, "work_id"})
)

var prometheusPusherStats *PrometheusPusherStats

func NewPrometheusPusherStats(gatewayURL, jobName string, mode string, reportName string) *PrometheusPusherStats {
	if prometheusPusherStats != nil {
		return prometheusPusherStats
	}
	prometheusPusherStats = &PrometheusPusherStats{
		pusher: push.New(gatewayURL, jobName).
			Grouping("instance", uuid.NewV1().String()).
			Grouping("mode", mode).
			Grouping("report", reportName),
	}
	return prometheusPusherStats
}

type PrometheusPusherStats struct {
	pusher *push.Pusher
}

func (p PrometheusPusherStats) OnStart() {
	resetPrometheusMetrics()
	log.Info().Msg("register prometheus metric collectors")
	registry := prometheus.NewRegistry()
	registry.MustRegister(masterGauge, workGauge)
	p.pusher = p.pusher.Gatherer(registry)
}

func (p PrometheusPusherStats) OnEvent(masterReport interfacecase.PerformanceReportMaster, workReports []interfacecase.PerformanceReportWork) {

	masterGauge.WithLabelValues(instanceMaster).Set(float64(masterReport.State))
	masterGauge.WithLabelValues(instanceMaster).Set(float64(masterReport.Workers))
	masterGauge.WithLabelValues(instanceMaster).Set(float64(masterReport.TargetUsers))
	masterGauge.WithLabelValues(instanceMaster).Set(float64(masterReport.CurrentUsers))

	for _, report := range workReports {
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(float64(report.State))
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(float64(report.Heartbeat))
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(float64(report.UserCount))
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(report.WorkerCpuUsage)
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(report.CpuUsage)
		if report.CpuWarningEmitted {
			workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(1)
		} else {
			workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(0)
		}
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(report.WorkerMemoryUsage)
		workGauge.WithLabelValues(instanceWorker, report.WorkID).Set(report.MemoryUsage)
	}

	if err := p.pusher.Push(); err != nil {
		log.Error().Err(err).Msg("push to Pushgateway failed")
	}
}

func (p PrometheusPusherStats) OnStop() {
	//TODO implement me
	panic("implement me")
}
