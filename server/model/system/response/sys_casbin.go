package response

import (
	"github.com/test-instructor/yangfan/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
