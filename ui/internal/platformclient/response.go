package platformclient

import (
	"encoding/json"
)

type APIResponse[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

func DecodeAPIResponse[T any](body []byte) (APIResponse[T], error) {
	var r APIResponse[T]
	err := json.Unmarshal(body, &r)
	return r, err
}
