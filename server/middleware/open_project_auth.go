package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/common/response"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	"gorm.io/gorm"
)

func OpenProjectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectId, ok, err := getRequestUint(c, "projectId")
		if err != nil {
			response.FailWithMessage("参数校验失败: "+err.Error(), c)
			c.Abort()
			return
		}
		if !ok || projectId == 0 {
			response.FailWithMessage("参数校验失败: projectId 不能为空", c)
			c.Abort()
			return
		}

		uuidStr, _, err := getRequestString(c, "uuid")
		if err != nil {
			response.FailWithMessage("参数校验失败: "+err.Error(), c)
			c.Abort()
			return
		}
		secretStr, _, err := getRequestString(c, "secret")
		if err != nil {
			response.FailWithMessage("参数校验失败: "+err.Error(), c)
			c.Abort()
			return
		}

		if uuidStr == "" {
			uuidStr, _, _ = getRequestString(c, "ci_uuid")
		}
		if secretStr == "" {
			secretStr, _, _ = getRequestString(c, "ci_secret")
		}

		if uuidStr == "" || secretStr == "" {
			response.FailWithMessage("参数校验失败: uuid/secret 不能为空", c)
			c.Abort()
			return
		}

		var pj projectmgr.Project
		if err := global.GVA_DB.First(&pj, projectId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.NoAuth("鉴权失败", c)
				c.Abort()
				return
			}
			response.FailWithMessage("鉴权失败: "+err.Error(), c)
			c.Abort()
			return
		}

		if pj.UUID != uuidStr || pj.Secret != secretStr {
			response.NoAuth("鉴权失败", c)
			c.Abort()
			return
		}

		c.Set("projectId", uint(projectId))
		c.Next()
	}
}

func getRequestString(c *gin.Context, name string) (string, bool, error) {
	if v := c.Query(name); v != "" {
		return v, true, nil
	}
	if !strings.HasPrefix(c.GetHeader("Content-Type"), "application/json") {
		return "", false, nil
	}
	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return "", false, nil
	}
	bodyBytes, err := c.GetRawData()
	if err != nil {
		return "", false, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if len(bodyBytes) == 0 {
		return "", false, nil
	}

	var raw map[string]any
	if err := json.Unmarshal(bodyBytes, &raw); err != nil {
		return "", false, nil
	}

	v, ok := raw[name]
	if !ok || v == nil {
		return "", false, nil
	}

	switch vv := v.(type) {
	case string:
		return vv, vv != "", nil
	case float64:
		return strconv.FormatInt(int64(vv), 10), true, nil
	default:
		return "", false, nil
	}
}

func getRequestUint(c *gin.Context, name string) (uint, bool, error) {
	if v := c.Query(name); v != "" {
		u, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, true, errors.New("字段 " + name + " 格式错误")
		}
		return uint(u), true, nil
	}

	s, ok, err := getRequestString(c, name)
	if err != nil {
		return 0, ok, err
	}
	if !ok || s == "" {
		return 0, false, nil
	}
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, true, errors.New("字段 " + name + " 格式错误")
	}
	return uint(u), true, nil
}
