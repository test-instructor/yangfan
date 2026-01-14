package runTestCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// DataWarehouseResponse represents the response from data warehouse API
type DataWarehouseResponse struct {
	Code  int                      `json:"code"`
	Count int                      `json:"count"`
	Total int                      `json:"total"`
	List  []map[string]interface{} `json:"list"`
	Msg   string                   `json:"msg"`
}

// queryDataWarehouse queries the data warehouse API and returns the data
func queryDataWarehouse(dataWarehouse datatypes.JSONMap, projectID int64, envID int) ([]map[string]interface{}, error) {
	if dataWarehouse == nil || len(dataWarehouse) == 0 {
		return nil, nil
	}
	// 如果未选择类型，直接跳过
	if t, ok := dataWarehouse["type"]; !ok || t == nil || (func(v interface{}) bool { s, _ := v.(string); return strings.TrimSpace(s) == "" })(t) {
		return nil, nil
	}

	// Add projectId and envId to dataWarehouse
	dataWarehouseBody := make(map[string]interface{})
	for k, v := range dataWarehouse {
		dataWarehouseBody[k] = v
	}
	dataWarehouseBody["projectId"] = projectID
	dataWarehouseBody["envId"] = envID

	// Marshal request body
	requestBody, err := json.Marshal(dataWarehouseBody)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal data warehouse request body", zap.Error(err))
		return nil, err
	}

	// Create HTTP request to data warehouse API
	dataWarehouseURL := fmt.Sprintf("http://%s:%d/api/datawarehouse/query", global.GVA_DW_HOST, global.GVA_DW_PORT)
	req, err := http.NewRequest("POST", dataWarehouseURL, bytes.NewReader(requestBody))
	if err != nil {
		global.GVA_LOG.Error("Failed to create data warehouse request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "YangFan-Client/V2")

	// Execute request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("Failed to query data warehouse", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("Failed to read data warehouse response", zap.Error(err))
		return nil, err
	}

	// Parse response
	var dwResp DataWarehouseResponse
	if err := json.Unmarshal(respBody, &dwResp); err != nil {
		global.GVA_LOG.Error("Failed to unmarshal data warehouse response", zap.Error(err))
		return nil, err
	}

	if dwResp.Code != 0 && dwResp.Code != 200 {
		global.GVA_LOG.Error("Data warehouse API returned error",
			zap.String("msg", dwResp.Msg),
			zap.Int("code", dwResp.Code))
		return nil, nil
	}

	// Check if count matches total
	if dwResp.Count != dwResp.Total {
		global.GVA_LOG.Warn("Data warehouse returned partial data",
			zap.Int("count", dwResp.Count),
			zap.Int("total", dwResp.Total))
	}

	global.GVA_LOG.Info("Data warehouse query successful",
		zap.Int("count", dwResp.Count),
		zap.Int("total", dwResp.Total),
		zap.Int("list_size", len(dwResp.List)))

	return dwResp.List, nil
}

// convertDataWarehouseToParameters converts data warehouse response to Parameters format
// Returns: parameters map, success flag
func convertDataWarehouseToParameters(data []map[string]interface{}) (map[string]interface{}, bool) {
	if len(data) == 0 {
		global.GVA_LOG.Warn("Data warehouse returned empty data")
		return nil, false
	}

	// Extract all column names from the first row
	var columnNames []string
	for key := range data[0] {
		columnNames = append(columnNames, key)
	}

	// Sort column names for consistency
	if len(columnNames) == 0 {
		global.GVA_LOG.Error("Data warehouse returned data with no columns")
		return nil, false
	}
	sort.Strings(columnNames)

	// Create combined key (e.g., "col1-col2-col3")
	combinedKey := strings.Join(columnNames, "-")

	// Build values array: each element is a row of values
	var values [][]interface{}
	for _, row := range data {
		var rowValues []interface{}
		for _, colName := range columnNames {
			rowValues = append(rowValues, row[colName])
		}
		values = append(values, rowValues)
	}

	// Create parameters map
	parameters := make(map[string]interface{})
	parameters[combinedKey] = values

	global.GVA_LOG.Info("Data warehouse data converted to parameters",
		zap.String("parameter_key", combinedKey),
		zap.Int("rows", len(values)))

	return parameters, true
}
