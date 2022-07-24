package response

import (
	"github.com/test-instructor/cheetah/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
