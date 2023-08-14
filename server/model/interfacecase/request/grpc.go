package request

type GrpcFunc struct {
	Host   *string `json:"host" form:"host"`
	Server *string `json:"server" form:"server"`
	Method *string `json:"method" form:"method"`
	Ref    *bool   `json:"ref" form:"ref"`
}
