package response

import "github.com/test-instructor/yangfan/server/v2/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
