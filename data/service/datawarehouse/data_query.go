package datawarehouse

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 查询操作符枚举
const (
	OpEQ         = "EQ"          // 等于
	OpNE         = "NE"          // 不等于
	OpGT         = "GT"          // 大于
	OpLT         = "LT"          // 小于
	OpGE         = "GE"          // 大于等于
	OpLE         = "LE"          // 小于等于
	OpBetween    = "BETWEEN"     // BETWEEN
	OpNotBetween = "NOT_BETWEEN" // NOT BETWEEN
	OpIsNull     = "IS_NULL"     // 为空
	OpIsNotNull  = "IS_NOT_NULL" // 不为空
	OpLike       = "LIKE"        // 模糊匹配
	OpNotLike    = "NOT_LIKE"    // 反向模糊匹配
)

// 逻辑运算类型
const (
	LogicAnd = "AND"
	LogicOr  = "OR"
)

// FieldCondition 单个字段的查询条件
type FieldCondition struct {
	Field    string      `json:"field"`            // JSON 字段名
	Operator string      `json:"operator"`         // 操作符: EQ, NE, GT, LT, GE, LE, BETWEEN, NOT_BETWEEN, IS_NULL, IS_NOT_NULL, LIKE, NOT_LIKE
	Value    interface{} `json:"value,omitempty"`  // 单值或范围值（BETWEEN时使用数组[v1,v2]）
	Value2   interface{} `json:"value2,omitempty"` // 范围结束值（可选，BETWEEN时也可用）
}

// ConditionGroup 条件组 - 包含多个字段条件
type ConditionGroup struct {
	Logic      string           `json:"logic,omitempty"` // AND / OR，组内所有条件的逻辑关系，默认AND，只有1个条件时可省略
	Conditions []FieldCondition `json:"conditions"`      // 筛选字段列表
}

// QueryFilter 查询筛选 - 包含多个条件组
type QueryFilter struct {
	Logic  string           `json:"logic,omitempty"` // AND / OR，所有条件组之间的逻辑关系，默认AND，只有1个组时可省略
	Groups []ConditionGroup `json:"groups"`          // 条件列表
}

// DataQueryRequest 数据查询请求
type DataQueryRequest struct {
	ProjectID int64        `json:"projectId"`        // 项目ID
	Type      string       `json:"type"`             // 数据类型
	EnvID     uint         `json:"envId"`            // 环境ID
	Status    *int         `json:"status,omitempty"` // 可选: 数据状态过滤
	Count     int          `json:"count"`            // 本次需要读取的数据条数
	Filter    *QueryFilter `json:"filter,omitempty"` // 可选的查询筛选条件
}

// DataQueryResponse 查询结果
type DataQueryResponse struct {
	Count int                      `json:"count"` // 请求的数据条数
	Total int                      `json:"total"` // 实际查询到的数据条数
	List  []map[string]interface{} `json:"list"`  // 只返回 value 字段的内容
}

type DataQueryService struct{}

// QueryData 根据条件查询数据，返回 value 字段内容和实际条数
// 说明：当查询的是「可用数据」(status=0 或未显式指定 status) 时，会在读取后将这些数据标记为「已占用」(status=1)。
func (s *DataQueryService) QueryData(req DataQueryRequest) ([]map[string]interface{}, int, error) {
	// 如果未显式指定 status，默认只读取可用数据
	requestedStatus := datawarehouse.DataStatusAvailable
	if req.Status != nil {
		requestedStatus = *req.Status
	}

	var list []datawarehouse.DataCategoryData

	// 使用事务保证：查询到的数据与随后的状态更新一致
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 构建基础查询
		db := tx.Model(&datawarehouse.DataCategoryData{}).
			Select("id", "value").
			Where("type = ? AND env_id = ? AND project_id = ?",
				req.Type, req.EnvID, req.ProjectID).
			Where("status = ?", requestedStatus).
			Order("id ASC")

		// 追加 Value 上的查询条件
		if req.Filter != nil {
			whereClause, args, err := buildFilterCondition(*req.Filter)
			if err != nil {
				return err
			}
			if whereClause != "" {
				db = db.Where(whereClause, args...)
			}
		}

		// 只有在“领取可用数据”时才需要加锁，避免并发下重复领取
		// 可选：MySQL 8+/PostgreSQL 可用 SKIP LOCKED 来避免并发下互相等待（通过环境变量开启）
		if requestedStatus == datawarehouse.DataStatusAvailable {
			lock := clause.Locking{Strength: "UPDATE"}
			if os.Getenv("DW_SKIP_LOCKED") == "true" {
				lock.Options = "SKIP LOCKED"
			}
			db = db.Clauses(lock)
		}

		if err := db.Limit(req.Count).Find(&list).Error; err != nil {
			return err
		}

		// 查询的不是可用数据，或没有查到数据，则不更新状态
		if requestedStatus != datawarehouse.DataStatusAvailable || len(list) == 0 {
			return nil
		}

		ids := make([]uint, 0, len(list))
		for _, item := range list {
			ids = append(ids, item.ID)
		}

		now := time.Now()
		// 读取成功后，将这些数据标记为已占用 (status=1)
		// 额外加上 status=0 的条件，避免误把已清洗(status=2)的数据改回已占用
		return tx.Model(&datawarehouse.DataCategoryData{}).
			Where("id IN ? AND status = ?", ids, datawarehouse.DataStatusAvailable).
			Updates(map[string]interface{}{
				"status":  datawarehouse.DataStatusUsed,
				"used_at": now,
			}).Error
	})
	if err != nil {
		return nil, 0, err
	}

	// 提取 value 字段内容 (datatypes.JSONMap 本身就是 map[string]interface{})
	result := make([]map[string]interface{}, 0, len(list))
	for _, item := range list {
		if item.Value != nil {
			result = append(result, map[string]interface{}(item.Value))
		}
	}

	return result, len(result), nil
}

// ValidateRequest 验证请求参数
func (s *DataQueryService) ValidateRequest(req *DataQueryRequest) error {
	if req.ProjectID == 0 || req.Type == "" || req.EnvID == 0 {
		return fmt.Errorf("projectId, type, envId 为必填项")
	}
	if req.Count <= 0 {
		req.Count = 50
	}
	if req.Count > 1000 {
		req.Count = 1000
	}
	return nil
}

// buildFilterCondition 将 QueryFilter 转换为 SQL 片段
func buildFilterCondition(filter QueryFilter) (string, []interface{}, error) {
	if len(filter.Groups) == 0 {
		return "", nil, nil
	}

	// 获取组间逻辑运算符，默认 AND
	filterLogic := strings.ToUpper(strings.TrimSpace(filter.Logic))
	if filterLogic == "" {
		filterLogic = LogicAnd
	}
	if filterLogic != LogicAnd && filterLogic != LogicOr {
		return "", nil, fmt.Errorf("不支持的组间逻辑运算符: %s", filter.Logic)
	}

	var (
		groupParts []string
		allArgs    []interface{}
	)

	// 遍历所有条件组
	for _, group := range filter.Groups {
		groupClause, groupArgs, err := buildConditionGroupClause(group)
		if err != nil {
			return "", nil, err
		}
		if groupClause == "" {
			continue
		}
		groupParts = append(groupParts, fmt.Sprintf("(%s)", groupClause))
		allArgs = append(allArgs, groupArgs...)
	}

	if len(groupParts) == 0 {
		return "", nil, nil
	}

	// 如果只有一个条件组，不需要外层括号
	if len(groupParts) == 1 {
		return groupParts[0][1 : len(groupParts[0])-1], allArgs, nil
	}

	joined := strings.Join(groupParts, fmt.Sprintf(" %s ", filterLogic))
	return joined, allArgs, nil
}

// buildConditionGroupClause 将 ConditionGroup 转换为 SQL 片段
func buildConditionGroupClause(group ConditionGroup) (string, []interface{}, error) {
	if len(group.Conditions) == 0 {
		return "", nil, nil
	}

	// 获取组内逻辑运算符，默认 AND
	groupLogic := strings.ToUpper(strings.TrimSpace(group.Logic))
	if groupLogic == "" {
		groupLogic = LogicAnd
	}
	if groupLogic != LogicAnd && groupLogic != LogicOr {
		return "", nil, fmt.Errorf("不支持的组内逻辑运算符: %s", group.Logic)
	}

	var (
		parts []string
		args  []interface{}
	)

	// 遍历所有字段条件
	for _, cond := range group.Conditions {
		clause, condArgs, err := buildFieldConditionClause(cond)
		if err != nil {
			return "", nil, err
		}
		if clause == "" {
			continue
		}
		parts = append(parts, clause)
		args = append(args, condArgs...)
	}

	if len(parts) == 0 {
		return "", nil, nil
	}

	// 如果只有一个字段条件，不需要逻辑运算符
	if len(parts) == 1 {
		return parts[0], args, nil
	}

	joined := strings.Join(parts, fmt.Sprintf(" %s ", groupLogic))
	return joined, args, nil
}

// buildFieldConditionClause 将单个 FieldCondition 转换为 SQL
func buildFieldConditionClause(cond FieldCondition) (string, []interface{}, error) {
	field := strings.TrimSpace(cond.Field)
	if field == "" {
		return "", nil, fmt.Errorf("字段名 field 不能为空")
	}

	op := strings.ToUpper(strings.TrimSpace(cond.Operator))
	if op == "" {
		op = OpEQ
	}

	jsonPath := fmt.Sprintf("$.%s", field)

	// 将值转换为 JSON 字符串，确保布尔值、字符串、数字都能正确比较
	toJSONString := func(v interface{}) (string, error) {
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	switch op {
	case OpEQ:
		jsonVal, err := toJSONString(cond.Value)
		if err != nil {
			return "", nil, err
		}
		return "JSON_EXTRACT(value, ?) = CAST(? AS JSON)", []interface{}{jsonPath, jsonVal}, nil
	case OpNE:
		jsonVal, err := toJSONString(cond.Value)
		if err != nil {
			return "", nil, err
		}
		return "JSON_EXTRACT(value, ?) <> CAST(? AS JSON)", []interface{}{jsonPath, jsonVal}, nil
	case OpGT:
		return "JSON_EXTRACT(value, ?) > ?", []interface{}{jsonPath, cond.Value}, nil
	case OpLT:
		return "JSON_EXTRACT(value, ?) < ?", []interface{}{jsonPath, cond.Value}, nil
	case OpGE:
		return "JSON_EXTRACT(value, ?) >= ?", []interface{}{jsonPath, cond.Value}, nil
	case OpLE:
		return "JSON_EXTRACT(value, ?) <= ?", []interface{}{jsonPath, cond.Value}, nil
	case OpBetween, OpNotBetween:
		v1, v2, err := extractBetweenValuesFromField(cond)
		if err != nil {
			return "", nil, err
		}
		if op == OpBetween {
			return "JSON_EXTRACT(value, ?) BETWEEN ? AND ?", []interface{}{jsonPath, v1, v2}, nil
		}
		return "JSON_EXTRACT(value, ?) NOT BETWEEN ? AND ?", []interface{}{jsonPath, v1, v2}, nil
	case OpIsNull:
		return "JSON_EXTRACT(value, ?) IS NULL", []interface{}{jsonPath}, nil
	case OpIsNotNull:
		return "JSON_EXTRACT(value, ?) IS NOT NULL", []interface{}{jsonPath}, nil
	case OpLike, OpNotLike:
		pattern := fmt.Sprintf("%v", cond.Value) // 用户自行控制通配符
		if op == OpLike {
			return "JSON_EXTRACT(value, ?) LIKE ?", []interface{}{jsonPath, pattern}, nil
		}
		return "JSON_EXTRACT(value, ?) NOT LIKE ?", []interface{}{jsonPath, pattern}, nil
	default:
		return "", nil, fmt.Errorf("不支持的操作符: %s", cond.Operator)
	}
}

// extractBetweenValuesFromField 解析 BETWEEN/NOT BETWEEN 的两个值
func extractBetweenValuesFromField(cond FieldCondition) (interface{}, interface{}, error) {
	if arr, ok := cond.Value.([]interface{}); ok {
		if len(arr) != 2 {
			return nil, nil, fmt.Errorf("BETWEEN 条件需要两个值")
		}
		return arr[0], arr[1], nil
	}

	if cond.Value == nil || cond.Value2 == nil {
		return nil, nil, fmt.Errorf("BETWEEN/NOT_BETWEEN 需要提供 value 和 value2")
	}
	return cond.Value, cond.Value2, nil
}

// 确保 gorm.DB 被使用（避免编译器警告）
var _ *gorm.DB
