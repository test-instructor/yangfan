package interfacecase

import (
	"gorm.io/gorm"
)

func projectDB(projectID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 添加其他参数对应的查询条件
		if projectID != 0 {
			db = db.Preload("Project").Joins("Project").Where("Project.ID = ?", projectID)
		}
		return db
	}
}

func menuDB(db *gorm.DB, menuid uint) *gorm.DB {
	return db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", menuid)
}

var _ = menuDB
