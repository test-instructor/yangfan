package datacategory

import (
	"fmt"
	"sync"

	"github.com/test-instructor/yangfan/data/internal/python"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"go.uber.org/zap"
)

// Service 数据分类处理服务
type Service struct{}

// NewService 创建服务实例
func NewService() *Service {
	return &Service{}
}

// ProcessAll 处理所有数据分类
// 主入口：定时任务调用此方法
func (s *Service) ProcessAll() {
	global.GVA_LOG.Info("开始处理所有数据分类")

	// 1. 查询所有数据分类
	categories, err := s.getAllCategories()
	if err != nil {
		global.GVA_LOG.Error("获取数据分类失败", zap.Error(err))
		return
	}

	if len(categories) == 0 {
		global.GVA_LOG.Info("没有数据分类需要处理")
		return
	}

	global.GVA_LOG.Info("获取到数据分类", zap.Int("count", len(categories)))

	// 2. 遍历处理每个数据分类
	var results []*ProcessResult
	for _, category := range categories {
		categoryResults := s.processCategory(&category)
		results = append(results, categoryResults...)
	}

	// 3. 输出处理结果汇总
	s.logSummary(results)
}

// getAllCategories 获取所有数据分类
func (s *Service) getAllCategories() ([]datawarehouse.DataCategoryManagement, error) {
	var categories []datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// processCategory 处理单个数据分类
// 同一 Category 下的所有环境并行处理
func (s *Service) processCategory(category *datawarehouse.DataCategoryManagement) []*ProcessResult {
	categoryName := ""
	if category.Name != nil {
		categoryName = *category.Name
	}

	global.GVA_LOG.Info("处理数据分类",
		zap.Uint("categoryId", category.ID),
		zap.String("categoryName", categoryName),
		zap.Int64("projectId", category.ProjectId),
	)

	// 获取项目下的所有环境
	envList, err := s.getProjectEnvs(category.ProjectId)
	if err != nil {
		global.GVA_LOG.Error("获取环境列表失败",
			zap.Uint("categoryId", category.ID),
			zap.Error(err),
		)
		return nil
	}

	if len(envList) == 0 {
		global.GVA_LOG.Warn("项目下没有环境",
			zap.Uint("categoryId", category.ID),
			zap.Int64("projectId", category.ProjectId),
		)
		return nil
	}

	// 用于收集所有环境的结果
	var (
		wg         sync.WaitGroup
		mu         sync.Mutex
		allResults []*ProcessResult
	)

	// 并行处理所有环境
	for _, env := range envList {
		wg.Add(1)
		go func(env platform.Env) {
			defer wg.Done()

			envResults := s.processEnv(category, categoryName, env)

			// 安全地收集结果
			mu.Lock()
			allResults = append(allResults, envResults...)
			mu.Unlock()
		}(env)
	}

	// 等待所有环境处理完成
	wg.Wait()

	global.GVA_LOG.Info("数据分类处理完成",
		zap.Uint("categoryId", category.ID),
		zap.String("categoryName", categoryName),
		zap.Int("envCount", len(envList)),
		zap.Int("resultCount", len(allResults)),
	)

	return allResults
}

// processEnv 处理单个环境
func (s *Service) processEnv(category *datawarehouse.DataCategoryManagement, categoryName string, env platform.Env) []*ProcessResult {
	var results []*ProcessResult

	ctx := &ExecuteContext{
		ProjectID:      category.ProjectId,
		CategoryID:     category.ID,
		CategoryName:   categoryName,
		Type:           category.Type,
		EnvID:          env.ID,
		EnvName:        env.Name,
		CreateCallType: category.CreateCallType,
		CleanCallType:  category.CleanCallType,
	}

	// 创建执行器
	executor := NewExecutor(ctx)

	// 前置操作：初始化 Python Runner（如果需要）
	var runner *python.Runner
	needPython := s.needPythonRunner(category)

	if needPython {
		runner = python.NewRunner(category.ProjectId, env.ID, category.ID)
		if err := runner.Init(); err != nil {
			global.GVA_LOG.Error("Python Runner 初始化失败",
				zap.Uint("categoryId", category.ID),
				zap.Uint("envId", env.ID),
				zap.Error(err),
			)
			// 初始化失败，返回失败结果
			results = append(results, &ProcessResult{
				CategoryID:   category.ID,
				CategoryName: categoryName,
				EnvID:        env.ID,
				Action:       "init",
				Success:      false,
				Message:      fmt.Sprintf("Python Runner 初始化失败: %v", err),
			})
			return results
		}
		// 确保后置操作：关闭 Runner
		defer runner.Close()

		// 设置 Runner 到 Executor
		executor.SetRunner(runner)
	}

	// 阶段一：执行清洗/更新
	cleanResult, err := executor.CleanData()
	if err != nil {
		global.GVA_LOG.Error("清洗数据失败",
			zap.Uint("categoryId", category.ID),
			zap.Uint("envId", env.ID),
			zap.Error(err),
		)
	}
	if cleanResult != nil {
		results = append(results, cleanResult)
	}

	// 阶段二：执行创建
	createResult, err := executor.CreateData()
	if err != nil {
		global.GVA_LOG.Error("创建数据失败",
			zap.Uint("categoryId", category.ID),
			zap.Uint("envId", env.ID),
			zap.Error(err),
		)
	}
	if createResult != nil {
		results = append(results, createResult)
	}

	return results
}

// needPythonRunner 判断是否需要 Python Runner
func (s *Service) needPythonRunner(category *datawarehouse.DataCategoryManagement) bool {
	// 如果 CleanCallType 或 CreateCallType 是 Python 类型，则需要 Runner
	if category.CleanCallType != nil && *category.CleanCallType == CallTypePython {
		return true
	}
	if category.CreateCallType != nil && *category.CreateCallType == CallTypePython {
		return true
	}
	return false
}

// getProjectEnvs 获取项目下的所有环境
func (s *Service) getProjectEnvs(projectID int64) ([]platform.Env, error) {
	var envList []platform.Env
	if err := global.GVA_DB.Where("project_id = ?", projectID).Find(&envList).Error; err != nil {
		return nil, fmt.Errorf("查询环境列表失败: %w", err)
	}
	return envList, nil
}

// logSummary 输出处理结果汇总
func (s *Service) logSummary(results []*ProcessResult) {
	if len(results) == 0 {
		global.GVA_LOG.Info("数据分类处理完成，无处理结果")
		return
	}

	successCount := 0
	failCount := 0
	cleanCount := 0
	createCount := 0

	for _, r := range results {
		if r.Success {
			successCount++
		} else {
			failCount++
		}
		if r.Action == "clean" {
			cleanCount += r.DataCount
		} else if r.Action == "create" {
			createCount += r.DataCount
		}
	}

	global.GVA_LOG.Info("数据分类处理完成",
		zap.Int("totalOperations", len(results)),
		zap.Int("successCount", successCount),
		zap.Int("failCount", failCount),
		zap.Int("cleanedDataCount", cleanCount),
		zap.Int("createdDataCount", createCount),
	)
}
