package utils

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/system"
	systemReq "github.com/test-instructor/yangfan/server/v2/model/system/request"
	"go.uber.org/zap"
)

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		j := NewJWT()
		token, _ = c.Cookie("x-token")
		if token == "" {
			return ""
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Warn("cookie token 解析失败，将清理 cookie x-token", zap.Error(err))
			ClearToken(c)
			return ""
		}
		if claims.ExpiresAt != nil {
			maxAgeSeconds := int(time.Until(claims.ExpiresAt.Time).Seconds())
			if maxAgeSeconds > 0 {
				SetToken(c, token, maxAgeSeconds)
			}
		}
	}
	return token
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if token != "" {
			global.GVA_LOG.Error("从 Gin 的 Context 获取 jwt 解析信息失败", zap.Error(err))
		}
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}

func LoginToken(user system.Login, node string, source string) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	authorityId := user.GetAuthorityId()
	projectId := user.GetProjectID()

	if source == "ui" {
		if u, ok := user.(*system.SysUser); ok {
			if u.UiAuthorityId != 0 {
				authorityId = u.UiAuthorityId
			}
			if u.UiProjectId != 0 {
				projectId = u.UiProjectId
			}
		}
	}

	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.GetUUID(),
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: authorityId,
		ProjectId:   projectId,
		Node:        node,
		Source:      source,
	})
	token, err = j.CreateToken(claims)
	return
}
