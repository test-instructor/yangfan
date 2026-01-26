package main

import (
	"embed"

	"yangfan-ui/internal/logger"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/zap"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize logger with defaults (will be reconfigured in app.startup)
	logger.Setup(logger.Config{
		Level:     "info",
		Prefix:    "[ https://github.com/test-instructor/yangfan/ui ]",
		Retention: 30,
	})
	logger.Info("Starting application...")

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "扬帆自动化测试平台-UI自动化节点",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		logger.Error("Application crashed", zap.Error(err))
	}
}
