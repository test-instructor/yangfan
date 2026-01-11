package utils

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	systemReq "github.com/test-instructor/yangfan/server/v2/model/system/request"
)

// GetProjectID 从Gin的Context中获取从jwt解析出来的项目ID
func GetProjectID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ProjectId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ProjectId
	}
}

func GetProjectIDInt64(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return int64(cl.BaseClaims.ProjectId)
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return int64(waitUse.BaseClaims.ProjectId)
	}
}

func GetProjectIDToString(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return "0"
		} else {
			return strconv.Itoa(int(cl.BaseClaims.ProjectId))
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return strconv.Itoa(int(waitUse.BaseClaims.ProjectId))
	}
}

// BytesToReadCloser 将 []byte 转换为 io.ReadCloser
func BytesToReadCloser(data []byte) io.ReadCloser {
	return io.NopCloser(strings.NewReader(string(data)))
}

// PeekRequestBody 读取请求体内容但不消耗它
func PeekRequestBody(c *gin.Context) ([]byte, error) {
	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return []byte{}, nil
	}

	// 读取请求体
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	// 恢复请求体
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}
