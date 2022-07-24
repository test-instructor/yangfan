package system

type SysUseProject struct {
	SysUserId uint   `gorm:"column:sys_user_id"`
	ProjectId string `gorm:"column:project_id"`
}

func (s *SysUseProject) TableName() string {
	return "sys_user_project"
}
