package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
)

type FileType uint

const (
	FileDebugTalk           FileType = 1
	FileDebugTalkDefault    FileType = 10
	FileDebugTalkGen        FileType = 2
	FileDebugTalkGenDefault FileType = 20
)

type ApiDebugTalk struct {
	global.GVA_MODEL
	Operator
	FileType  FileType       `json:"file_type" gorm:"comment:文件类型"`
	Content   string         `json:"content" form:"content" gorm:"column:content;type:text;comment:文件内容"`
	ProjectID uint           `json:"-" gorm:"comment:所属项目"`
	Project   system.Project `json:"-"`
}
