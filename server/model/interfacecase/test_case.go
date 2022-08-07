// 自动生成模板TestCase
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/system"
)

// ApiTestCase 结构体
// 如果含有time.Time 请自行import time包
type ApiTestCase struct {
	global.GVA_MODEL
	Name      string         `json:"name" form:"name" gorm:"column:name;comment:;"`
	FrontCase *bool          `json:"front_case" orm:"front_case"`
	TStep     []ApiStep      `json:"TStep" form:"TStep" gorm:"many2many:ApiCaseRelationship;"`
	TimerTask []TimerTask    `json:"case" form:"case" gorm:"many2many:TimerTaskRelationship;"`
	ProjectID uint           `json:"-"`
	Project   system.Project `json:"-"`
	ApiMenuID uint           `json:"-"`
	ApiMenu   ApiMenu        `json:"-"`
}
