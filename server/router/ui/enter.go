package ui

import api "github.com/test-instructor/yangfan/server/v2/api/v1"

type RouterGroup struct {
	UINodeMenuRouter
}

var (
	uiNodeMenuApi = api.ApiGroupApp.UiApiGroup.UINodeMenuApi
)
