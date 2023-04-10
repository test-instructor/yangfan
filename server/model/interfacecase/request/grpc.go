package request

type GrpcFunc struct {
	Host   *string `json:"host"`
	Server *string `json:"server"`
	Method *string `json:"method"`
	Ref    *bool   `json:"ref"`
}
