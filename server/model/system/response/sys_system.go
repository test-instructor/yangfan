package response

import "github.com/test-instructor/yangfan/server/v2/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
