package response

import "github.com/test-instructor/yangfan/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
