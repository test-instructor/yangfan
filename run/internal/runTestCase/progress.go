package runTestCase

import (
	"context"
	"fmt"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"go.uber.org/zap"
)

// Redis key prefix and TTL settings for test report progress
const (
	RedisReportKeyPrefix = "test_report_progress"
	// Default TTL for progress keys; aligns with requirement of keeping
	// progress data for a limited time after task completion.
	RedisReportTTL = 24 * time.Hour
)

// ReportProgressTotals holds pre-calculated totals for a single report.
type ReportProgressTotals struct {
	TotalCases int
	TotalSteps int
	TotalApis  int
}

// BuildReportProgressKey formats a redis key like: test_report_progress:{report_id}:{field}
// Exported for use in AutoReport.LoadProgressFromRedis
func BuildReportProgressKey(reportID uint, field string) string {
	return fmt.Sprintf("%s:%d:%s", RedisReportKeyPrefix, reportID, field)
}

// initReportProgress initializes total counters and executed counters (0).
// If Redis is available, only writes to Redis (database will be updated on finalize).
// If Redis is not available, writes to database directly.
// Returns 0 since ProgressID is not created until finalize when using Redis.
func initReportProgress(reportID uint, totals ReportProgressTotals) uint {
	if reportID == 0 {
		return 0
	}

	// 优先使用 Redis
	if global.GVA_REDIS != nil {
		initReportProgressRedis(reportID, totals)
		return 0 // Redis 模式下不立即创建数据库记录
	}

	// Redis 不可用，回退到数据库
	return initReportProgressDB(reportID, totals)
}

// initReportProgressRedis 初始化进度到 Redis
func initReportProgressRedis(reportID uint, totals ReportProgressTotals) {
	ctx := context.Background()
	pipe := global.GVA_REDIS.TxPipeline()

	// Totals
	pipe.Set(ctx, BuildReportProgressKey(reportID, "total_cases"), totals.TotalCases, RedisReportTTL)
	pipe.Set(ctx, BuildReportProgressKey(reportID, "total_steps"), totals.TotalSteps, RedisReportTTL)
	pipe.Set(ctx, BuildReportProgressKey(reportID, "total_apis"), totals.TotalApis, RedisReportTTL)

	// Executed counters start from 0
	pipe.Set(ctx, BuildReportProgressKey(reportID, "executed_cases"), 0, RedisReportTTL)
	pipe.Set(ctx, BuildReportProgressKey(reportID, "executed_steps"), 0, RedisReportTTL)
	pipe.Set(ctx, BuildReportProgressKey(reportID, "executed_apis"), 0, RedisReportTTL)

	if _, err := pipe.Exec(ctx); err != nil {
		global.GVA_LOG.Error("init test report progress redis failed", zap.Uint("report_id", reportID), zap.Error(err))
	}
}

// initReportProgressDB 初始化报告进度到数据库
// 返回 ProgressID（失败时返回 0）
func initReportProgressDB(reportID uint, totals ReportProgressTotals) uint {
	if global.GVA_DB == nil {
		return 0
	}

	// 创建 Progress 记录
	progress := automation.AutoReportProgress{
		TotalCases:    totals.TotalCases,
		TotalSteps:    totals.TotalSteps,
		TotalApis:     totals.TotalApis,
		ExecutedCases: 0,
		ExecutedSteps: 0,
		ExecutedApis:  0,
	}

	// 检查报告是否已有 ProgressID
	var report automation.AutoReport
	if err := global.GVA_DB.Select("progress_id").First(&report, "id = ?", reportID).Error; err != nil {
		global.GVA_LOG.Error("查询报告失败", zap.Uint("report_id", reportID), zap.Error(err))
		return 0
	}

	if report.ProgressID != nil && *report.ProgressID != 0 {
		// 已存在 Progress，更新
		err := global.GVA_DB.Model(&automation.AutoReportProgress{}).Where("id = ?", *report.ProgressID).Updates(map[string]interface{}{
			"total_cases":    totals.TotalCases,
			"total_steps":    totals.TotalSteps,
			"total_apis":     totals.TotalApis,
			"executed_cases": 0,
			"executed_steps": 0,
			"executed_apis":  0,
		}).Error
		if err != nil {
			global.GVA_LOG.Error("更新报告进度失败", zap.Uint("report_id", reportID), zap.Error(err))
		}
		return *report.ProgressID
	}

	// 创建新的 Progress 记录
	if err := global.GVA_DB.Create(&progress).Error; err != nil {
		global.GVA_LOG.Error("创建报告进度失败", zap.Uint("report_id", reportID), zap.Error(err))
		return 0
	}

	// 更新 AutoReport 的 ProgressID
	if err := global.GVA_DB.Model(&automation.AutoReport{}).Where("id = ?", reportID).Update("progress_id", progress.ID).Error; err != nil {
		global.GVA_LOG.Error("更新报告 ProgressID 失败", zap.Uint("report_id", reportID), zap.Error(err))
	}
	return progress.ID
}

// incrReportProgress increments a specific executed counter atomically.
// field should be one of: case_executed, step_executed, api_executed.
// If Redis is available, only updates Redis. Otherwise updates database.
func incrReportProgress(reportID uint, field string, delta int64) {
	if reportID == 0 {
		return
	}

	// 优先使用 Redis
	if global.GVA_REDIS != nil {
		ctx := context.Background()
		// 转换字段名为 Redis key 格式
		redisField := convertFieldToRedisKey(field)
		key := BuildReportProgressKey(reportID, redisField)
		if err := global.GVA_REDIS.IncrBy(ctx, key, delta).Err(); err != nil {
			global.GVA_LOG.Error("incr test report progress redis failed", zap.Uint("report_id", reportID), zap.String("field", field), zap.Error(err))
		}
		return
	}

	// Redis 不可用，回退到数据库
	incrReportProgressDB(reportID, field, int(delta))
}

// convertFieldToRedisKey 将字段名转换为 Redis key 格式
func convertFieldToRedisKey(field string) string {
	switch field {
	case "case_executed":
		return "executed_cases"
	case "step_executed":
		return "executed_steps"
	case "api_executed":
		return "executed_apis"
	default:
		return field
	}
}

// incrReportProgressDB 更新数据库中的已执行数量
func incrReportProgressDB(reportID uint, field string, delta int) {
	if global.GVA_DB == nil {
		return
	}

	column := convertFieldToRedisKey(field)
	if column == field && field != "executed_cases" && field != "executed_steps" && field != "executed_apis" {
		return // 无效字段
	}

	// 通过 AutoReport 的 ProgressID 查找 Progress 并更新
	var report automation.AutoReport
	if err := global.GVA_DB.Select("progress_id").First(&report, "id = ?", reportID).Error; err != nil {
		global.GVA_LOG.Error("查询报告失败", zap.Uint("report_id", reportID), zap.Error(err))
		return
	}

	if report.ProgressID == nil || *report.ProgressID == 0 {
		global.GVA_LOG.Warn("报告无关联的进度记录", zap.Uint("report_id", reportID))
		return
	}

	err := global.GVA_DB.Model(&automation.AutoReportProgress{}).
		Where("id = ?", *report.ProgressID).
		Update(column, global.GVA_DB.Raw(fmt.Sprintf("%s + ?", column), delta)).Error
	if err != nil {
		global.GVA_LOG.Error("更新报告进度到数据库失败", zap.Uint("report_id", reportID), zap.String("field", field), zap.Error(err))
	}
}

// finalizeReportProgress syncs Redis progress data to database and sets up associations.
// This should be called when test execution completes.
func finalizeReportProgress(reportID uint) {
	if reportID == 0 {
		return
	}

	// 如果 Redis 不可用，数据已经在数据库中，无需操作
	if global.GVA_REDIS == nil {
		return
	}

	// 从 Redis 读取进度数据
	ctx := context.Background()
	pipe := global.GVA_REDIS.TxPipeline()

	totalCasesCmd := pipe.Get(ctx, BuildReportProgressKey(reportID, "total_cases"))
	totalStepsCmd := pipe.Get(ctx, BuildReportProgressKey(reportID, "total_steps"))
	totalApisCmd := pipe.Get(ctx, BuildReportProgressKey(reportID, "total_apis"))

	if _, err := pipe.Exec(ctx); err != nil {
		global.GVA_LOG.Error("read progress from redis failed", zap.Uint("report_id", reportID), zap.Error(err))
		return
	}

	// 解析数据
	totalCases, _ := totalCasesCmd.Int()
	totalSteps, _ := totalStepsCmd.Int()
	totalApis, _ := totalApisCmd.Int()

	// 强制设置进度为 100% (无论是否有失败跳过)
	executedCases := totalCases
	executedSteps := totalSteps
	executedApis := totalApis

	// 写入数据库
	if global.GVA_DB != nil {
		progress := automation.AutoReportProgress{
			TotalCases:    totalCases,
			TotalSteps:    totalSteps,
			TotalApis:     totalApis,
			ExecutedCases: executedCases,
			ExecutedSteps: executedSteps,
			ExecutedApis:  executedApis,
		}

		// 检查报告是否已有 ProgressID
		var report automation.AutoReport
		if err := global.GVA_DB.Select("progress_id").First(&report, "id = ?", reportID).Error; err != nil {
			global.GVA_LOG.Error("查询报告失败", zap.Uint("report_id", reportID), zap.Error(err))
			return
		}

		if report.ProgressID != nil && *report.ProgressID != 0 {
			// 已存在 Progress，更新
			err := global.GVA_DB.Model(&automation.AutoReportProgress{}).Where("id = ?", *report.ProgressID).Updates(map[string]interface{}{
				"total_cases":    totalCases,
				"total_steps":    totalSteps,
				"total_apis":     totalApis,
				"executed_cases": executedCases,
				"executed_steps": executedSteps,
				"executed_apis":  executedApis,
			}).Error
			if err != nil {
				global.GVA_LOG.Error("更新报告进度失败", zap.Uint("report_id", reportID), zap.Error(err))
			}
		} else {
			// 创建新的 Progress 记录
			if err := global.GVA_DB.Create(&progress).Error; err != nil {
				global.GVA_LOG.Error("创建报告进度失败", zap.Uint("report_id", reportID), zap.Error(err))
				return
			}

			// 更新 AutoReport 的 ProgressID
			if err := global.GVA_DB.Model(&automation.AutoReport{}).Where("id = ?", reportID).Update("progress_id", progress.ID).Error; err != nil {
				global.GVA_LOG.Error("更新报告 ProgressID 失败", zap.Uint("report_id", reportID), zap.Error(err))
			}
		}
	}

	// 刷新 Redis TTL，保留一段时间以便查询
	keys := []string{
		BuildReportProgressKey(reportID, "total_cases"),
		BuildReportProgressKey(reportID, "executed_cases"),
		BuildReportProgressKey(reportID, "total_steps"),
		BuildReportProgressKey(reportID, "executed_steps"),
		BuildReportProgressKey(reportID, "total_apis"),
		BuildReportProgressKey(reportID, "executed_apis"),
	}

	pipe = global.GVA_REDIS.TxPipeline()
	for _, key := range keys {
		pipe.Expire(ctx, key, RedisReportTTL)
	}

	if _, err := pipe.Exec(ctx); err != nil {
		global.GVA_LOG.Error("finalize test report progress redis failed", zap.Uint("report_id", reportID), zap.Error(err))
	}
}
