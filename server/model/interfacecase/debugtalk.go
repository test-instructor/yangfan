package interfacecase

import (
	"github.com/test-instructor/yangfan/server/global"
)

type FileType uint

const (
	FileDebugTalk           FileType = 1
	FileDebugTalkDefault    FileType = 10
	FileDebugTalkGen        FileType = 2
	FileDebugTalkGenDefault FileType = 20
)

var _ = []FileType{FileDebugTalkDefault, FileDebugTalkGen, FileDebugTalkGenDefault}

type ApiDebugTalk struct {
	global.GVA_MODEL
	Operator
	FileType FileType `json:"file_type" gorm:"comment:文件类型"`
	Content  string   `json:"content,omitempty" form:"content" gorm:"column:content;type:text;comment:文件内容"`
}
