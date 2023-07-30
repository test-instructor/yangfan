package system

type SysUseProject struct {
	SysUser   SysUser
	SysUserId uint `gorm:"column:sys_user_id"`
	ProjectId uint `gorm:"column:project_id"`
	Select    bool `json:"select"`
	Delete    bool `json:"delete"`
	Save      bool `json:"save"`
	//Update    bool   `json:"update"`
	//Create    bool   `json:"create"`
}

func (s *SysUseProject) TableName() string {
	return "sys_user_project"
}
