package ui

import "github.com/test-instructor/yangfan/server/v2/service"

type ApiGroup struct {
	UINodeMenuApi
}

var (
	uiNodeMenuService = service.ServiceGroupApp.UiServiceGroup.UINodeMenuService
)
