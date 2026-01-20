package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/test-instructor/yangfan/data/router"
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

// Server 封装 HTTP 服务
// 这里只提供一个无需鉴权的数据查询接口，监听在固定端口上
// 后续如有需要可以在此扩展更多路由组
//
// NOTE: 为了避免与主服务的路由、中间件耦合，这里使用独立的 gin.Engine
// 仅依赖全局的日志和数据库实例。
type Server struct {
	engine *gin.Engine
	http   *http.Server
}

// NewServer 创建 HTTP Server 实例
func NewServer() *Server {
	// 使用默认的 gin 引擎，但不挂载任何鉴权中间件
	// 只开启最基础的日志和恢复中间件
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	s := &Server{engine: engine}
	s.registerRoutes()
	return s
}

// Start 启动 HTTP 服务（阻塞当前 goroutine）
func (s *Server) Start(addr string) error {
	s.http = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	global.GVA_LOG.Info("HTTP 服务启动", zap.String("addr", addr))

	// ListenAndServe 在正常 Shutdown 时会返回 http.ErrServerClosed
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.GVA_LOG.Error("HTTP 服务异常退出", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("HTTP 服务已关闭")
	return nil
}

// Shutdown 优雅关闭 HTTP 服务
func (s *Server) Shutdown(ctx context.Context) error {
	if s == nil || s.http == nil {
		return nil
	}
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.http.Shutdown(shutdownCtx)
}

// registerRoutes 注册所有路由
func (s *Server) registerRoutes() {
	// 公共路由组（无需鉴权）
	publicGroup := s.engine.Group("/api")
	publicGroup.Use(func(c *gin.Context) {
		if global.GVA_DB == nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"message": "database not ready"})
			return
		}
		c.Next()
	})

	// 注册业务路由
	initBizRouter(publicGroup)
}

// initBizRouter 初始化业务路由
// 后续添加新路由组可以在此扩展
func initBizRouter(publicGroup *gin.RouterGroup) {
	// Datawarehouse 路由组
	datawarehouseRouter := router.RouterGroupApp.Datawarehouse
	datawarehouseRouter.InitDataQueryRouter(publicGroup)

	// 后续可以添加更多路由组，例如：
	// otherRouter := router.RouterGroupApp.Other
	// otherRouter.InitOtherRouter(publicGroup)
}
