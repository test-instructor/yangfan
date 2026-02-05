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
	type rawAPIResponse struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
		Msg  string          `json:"msg"`
	}

	var raw rawAPIResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return APIResponse[T]{}, err
	}

	r := APIResponse[T]{Code: raw.Code, Msg: raw.Msg}
	if raw.Code != 0 {
		return r, nil
	}

	if len(raw.Data) == 0 || string(raw.Data) == "null" {
		var zero T
		r.Data = zero
		return r, nil
	}

	if err := json.Unmarshal(raw.Data, &r.Data); err != nil {
		return APIResponse[T]{}, err
	}
	return r, nil
}
