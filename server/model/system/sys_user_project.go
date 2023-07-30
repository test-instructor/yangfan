package system

type SysUseProject struct {
	SysUser   SysUser
	SysUserId uint   `gorm:"column:sys_user_id"`
	ProjectId string `gorm:"column:project_id"`
	Read      bool   `json:"read"`
	Delete    bool   `json:"delete"`
	Update    bool   `json:"update"`
	Create    bool   `json:"create"`
}

func (s *SysUseProject) TableName() string {
	return "sys_user_project"
}
