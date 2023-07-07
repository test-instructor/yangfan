package boomer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	uuid "github.com/satori/go.uuid"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"go.uber.org/zap"
)

var (
	instanceMaster = "master"
	instanceWorker = "worker"
)

var (
	stateGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "master_state",
		Help: "Performance report for master state",
	})

	workersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "master_workers",
		Help: "Performance report for master workers",
	})

	targetUsersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "master_target_users",
		Help: "Performance report for master target users",
	})

	currentUsersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "master_current_users",
		Help: "Performance report for master current users",
	})
)

var (
	stateGaugeWork = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_state",
		Help: "Performance report for work state",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})

	heartbeatGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_heartbeat",
		Help: "Performance report for work heartbeat",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})

	userCountGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_user_count",
		Help: "Performance report for work user count",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})

	workerCpuUsageGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_worker_cpu_usage",
		Help: "Performance report for work worker_cpu_usage",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})

	cpuUsageGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_cpu_usage",
		Help: "Performance report for work cpu usage",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})
	workerMemoryUsageGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_worker_memory_usage",
		Help: "Performance report for worker_memory usage",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})

	memoryUsageGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "work_memory_usage",
		Help: "Performance report for memory usage",
	}, []string{instanceWorker, "work_id", "ip", "os", "arch"})
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
	global.GVA_LOG.Info("register prometheus metric collectors")
	registry := prometheus.NewRegistry()
	registry.MustRegister(
		stateGauge,
		workersGauge,
		targetUsersGauge,
		currentUsersGauge,
		stateGaugeWork,
		heartbeatGauge,
		userCountGauge,
		workerCpuUsageGauge,
		cpuUsageGauge,
		workerMemoryUsageGauge,
		memoryUsageGauge,
	)
	p.pusher = p.pusher.Gatherer(registry)
}

func (p PrometheusPusherStats) OnEvent(masterReport interfacecase.PerformanceReportMaster, workReports []interfacecase.PerformanceReportWork) {

	stateGauge.Set(float64(masterReport.State))
	workersGauge.Set(float64(masterReport.Workers))
	targetUsersGauge.Set(float64(masterReport.TargetUsers))
	currentUsersGauge.Set(float64(masterReport.CurrentUsers))

	for _, worker := range workReports {
		userCountGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(float64(worker.UserCount))
		stateGaugeWork.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(float64(worker.State))
		heartbeatGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(float64(worker.Heartbeat))
		workerCpuUsageGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(worker.WorkerCpuUsage)
		cpuUsageGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(worker.CpuUsage)
		workerMemoryUsageGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(worker.WorkerMemoryUsage)
		memoryUsageGauge.WithLabelValues(instanceWorker, worker.WorkID, worker.IP, worker.OS, worker.Arch).Set(worker.MemoryUsage)

	}

	if err := p.pusher.Push(); err != nil {
		global.GVA_LOG.Error("push to Pushgateway failed", zap.Error(err))
	}
}

func (p PrometheusPusherStats) OnStop() {
	//TODO implement me
	panic("implement me")
}
