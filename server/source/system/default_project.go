package system

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/v2/model/projectmgr"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"github.com/test-instructor/yangfan/server/v2/service/system"
	"gorm.io/gorm"
)

const initOrderDefaultProject = initOrderUser + 1

type initDefaultProject struct{}

func init() {
	system.RegisterInit(initOrderDefaultProject, &initDefaultProject{})
}

func (i *initDefaultProject) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&projectmgr.Project{}, &projectmgr.UserProjectAccess{})
}

func (i *initDefaultProject) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&projectmgr.Project{}) && db.Migrator().HasTable(&projectmgr.UserProjectAccess{})
}

func (i *initDefaultProject) InitializerName() string {
	return "default_project_and_admin_access"
}

func (i *initDefaultProject) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	const defaultProjectUUID = "default"
	const defaultProjectName = "默认项目"

	err = db.Transaction(func(tx *gorm.DB) error {
		var admin sysModel.SysUser
		if err := tx.Where("username = ?", "admin").First(&admin).Error; err != nil {
			return errors.Wrap(err, "查找 admin 用户失败")
		}

		var pj projectmgr.Project
		pjQuery := tx.Where("uuid = ?", defaultProjectUUID).Limit(1).Find(&pj)
		if pjQuery.Error != nil {
			return errors.Wrap(pjQuery.Error, "查询默认项目失败")
		}
		if pjQuery.RowsAffected == 0 {
			pj = projectmgr.Project{
				Name:     defaultProjectName,
				Admin:    admin.ID,
				Creator:  admin.ID,
				Describe: "系统初始化项目",
				UUID:     defaultProjectUUID,
				Secret:   uuid.NewString(),
			}
			if err := tx.Create(&pj).Error; err != nil {
				return errors.Wrap(err, "创建默认项目失败")
			}
		}

		var upa projectmgr.UserProjectAccess
		upaQuery := tx.Where("user_id = ? AND project_id = ?", admin.ID, pj.ID).Limit(1).Find(&upa)
		if upaQuery.Error != nil {
			return errors.Wrap(upaQuery.Error, "查询 admin 项目权限失败")
		}
		if upaQuery.RowsAffected == 0 {
			upa = projectmgr.UserProjectAccess{
				UserId:           admin.ID,
				ProjectId:        pj.ID,
				AccessPermission: true,
				EditPermission:   true,
				DeletePermission: true,
			}
			if err := tx.Create(&upa).Error; err != nil {
				return errors.Wrap(err, "写入 admin 项目权限失败")
			}
		} else {
			updates := map[string]any{
				"access_permission": true,
				"edit_permission":   true,
				"delete_permission": true,
			}
			if err := tx.Model(&projectmgr.UserProjectAccess{}).Where("id = ?", upa.ID).Updates(updates).Error; err != nil {
				return errors.Wrap(err, "更新 admin 项目权限失败")
			}
		}

		if admin.ProjectID != pj.ID {
			if err := tx.Model(&sysModel.SysUser{}).Where("id = ?", admin.ID).Update("project_id", pj.ID).Error; err != nil {
				return errors.Wrap(err, "更新 admin 当前项目失败")
			}
		}

		return nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (i *initDefaultProject) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	var admin sysModel.SysUser
	if err := db.Where("username = ?", "admin").First(&admin).Error; err != nil {
		return false
	}

	var pj projectmgr.Project
	pjQuery := db.Where("uuid = ?", "default").Limit(1).Find(&pj)
	if pjQuery.Error != nil || pjQuery.RowsAffected == 0 {
		return false
	}

	var upa projectmgr.UserProjectAccess
	upaQuery := db.Where("user_id = ? AND project_id = ?", admin.ID, pj.ID).Limit(1).Find(&upa)
	if upaQuery.Error != nil || upaQuery.RowsAffected == 0 {
		return false
	}
	return upa.AccessPermission && upa.EditPermission && upa.DeletePermission
}
