package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	"github.com/test-instructor/yangfan/server/v2/utils"
	"gorm.io/gorm"
)

// ProjectAuth 项目权限验证中间件
func ProjectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从权限token中获取项目ID
		tokenProjectId := utils.GetProjectID(c)
		userId := utils.GetUserID(c)

		// 验证项目权限
		if !hasProjectPermission(c, userId, tokenProjectId) {
			response.NoAuth("没有项目访问权限", c)
			c.Abort()
			return
		}

		// 根据请求方法从不同位置获取请求中的项目ID
		var requestProjectId uint
		var err error

		switch c.Request.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodDelete:
			// 从URL参数中获取projectId
			requestProjectId, err = getProjectIdFromQuery(c)
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			// POST、PUT、DELETE请求从请求体中获取projectId
			requestProjectId, err = getProjectIdFromBody(c)
		default:
			response.FailWithMessage("不支持的请求方法", c)
			c.Abort()
			return
		}

		if err != nil {
			response.FailWithMessage("获取项目ID失败: "+err.Error(), c)
			c.Abort()
			return
		}

		//// 如果请求中没有项目ID，使用token中的项目ID
		//if requestProjectId == 0 {
		//	requestProjectId = tokenProjectId
		//}

		// 对比token中的项目ID和请求中的项目ID
		if tokenProjectId != requestProjectId {
			response.FailWithMessage("请勿操作其他项目的数据", c)
			c.Abort()
			return
		}

		// 将项目ID设置到上下文中，供后续处理使用
		c.Set("projectId", tokenProjectId)
		c.Next()
	}
}

// getProjectIdFromQuery 从URL查询参数中获取项目ID
func getProjectIdFromQuery(c *gin.Context) (uint, error) {
	projectIdStr := c.Query("projectId")
	if projectIdStr == "" {
		// 如果没有projectId参数，返回0不算错误
		return 0, errors.New("你没有权限操作该项目")
	}

	projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
	if err != nil {
		return 0, errors.New("项目ID格式错误")
	}

	return uint(projectId), nil
}

// ProjectIdBody 用于解析请求体中的项目ID
type ProjectIdBody struct {
	ProjectId uint `json:"projectId"`
}

// hasProjectPermission 检查用户对项目的权限
func hasProjectPermission(c *gin.Context, userId uint, projectId uint) bool {
	var userProjectAccess projectmgr.UserProjectAccess

	// 查询用户项目权限记录
	err := global.GVA_DB.Where("user_id = ? AND project_id = ?", userId, projectId).
		First(&userProjectAccess).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		// 数据库查询错误，默认拒绝访问
		return false
	}

	// 根据请求方法检查对应权限
	method := c.Request.Method
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		// 读取操作需要访问权限
		return userProjectAccess.AccessPermission
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		// 写入操作需要编辑权限
		return userProjectAccess.EditPermission
	case http.MethodDelete:
		// 删除操作需要删除权限
		return userProjectAccess.DeletePermission
	default:
		// 其他方法默认需要访问权限
		return userProjectAccess.AccessPermission
	}
}

// getProjectIdFromBody 从请求体中获取项目ID
func getProjectIdFromBody(c *gin.Context) (uint, error) {
	// 方法1：尝试从URL参数中获取
	if projectIdStr := c.Query("projectId"); projectIdStr != "" {
		projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
		if err != nil {
			return 0, errors.New("项目ID格式错误")
		}
		return uint(projectId), nil
	}

	// 方法2：检查Content-Type
	contentType := c.GetHeader("Content-Type")
	if contentType != "application/json" {
		return 0, nil // 非JSON请求，跳过body解析
	}

	// 方法3：使用c.GetRawData()并恢复
	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return 0, nil
	}

	// 读取原始数据
	bodyBytes, err := c.GetRawData()
	if err != nil {
		return 0, err
	}

	// 立即恢复请求体
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if len(bodyBytes) == 0 {
		return 0, nil
	}

	// 解析JSON
	var body struct {
		ProjectId *uint `json:"projectId"`
	}

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		// JSON解析失败，可能不是JSON格式或者是部分数据
		return 0, nil
	}

	if body.ProjectId != nil {
		return *body.ProjectId, nil
	}

	return 0, nil
}
