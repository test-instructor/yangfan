package runTestCase

import (
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type ApiCaseStep struct {
	ID        uint
	Name      string                  `json:"name" form:"name" gorm:"column:name;comment:;"`
	FrontCase *bool                   `json:"front_case" orm:"front_case"`
	TStep     []interfacecase.ApiStep `json:"TStep" form:"TStep" gorm:"many2many:ApiCaseStepRelationship;"`
	ApiCase   []interfacecase.ApiCase `json:"case" form:"case" gorm:"many2many:ApiCaseRelationship;"`
	ProjectID uint                    `json:"-"`
	Config    *interfacecase.ApiConfig
}
