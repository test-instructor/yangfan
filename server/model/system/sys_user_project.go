package system

type SysUserProject struct {
	SysUser   *SysUser `json:"-"`
	SysUserID uint     `gorm:"column:sys_user_id"`
	Project   *Project `json:"-"`
	ProjectID uint     `gorm:"column:project_id"`
	Select    bool     `json:"select" gorm:"column:select"`
	Delete    bool     `json:"delete" gorm:"column:delete"`
	Save      bool     `json:"save" gorm:"column:save"`
	Username  string   `json:"username" gorm:"-"`
	//Update    bool   `json:"update"`
	//Create    bool   `json:"create"`
}

func (s *SysUserProject) TableName() string {
	return "sys_user_project"
}
