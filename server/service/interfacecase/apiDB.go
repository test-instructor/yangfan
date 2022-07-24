package interfacecase

import (
	"gorm.io/gorm"
)

func projectDB(db *gorm.DB, projectid uint) *gorm.DB {
	return db.Preload("Project").Joins("Project").Where("Project.ID = ?", projectid)
}

func menuDB(db *gorm.DB, menuid uint) *gorm.DB {
	return db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", menuid)
}
