package response

import "github.com/test-instructor/cheetah/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
