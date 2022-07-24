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
	FileType    FileType `json:"file_type"`
	Content     string   `json:"content" form:"content" gorm:"column:content;type:text"`
	CreatedBy   system.SysUser
	CreatedByID uint
	ProjectID   uint           `json:"-"`
	Project     system.Project `json:"-"`
}
