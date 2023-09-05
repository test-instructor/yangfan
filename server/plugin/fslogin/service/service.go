package service

import (
	"errors"

	global2 "github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/system"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/global"
	"github.com/test-instructor/yangfan/server/plugin/fslogin/model"
	"github.com/test-instructor/yangfan/server/service"
	"gorm.io/gorm"
)

type FsLoginService struct{}

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func (e *FsLoginService) LoginOrRegister(FsUserInfo model.FsUserInfo) (gvaInfo system.SysUser, err error) {
	var loginUserInfo model.FsUserInfo
	err = global2.GVA_DB.Where("union_id = ?", FsUserInfo.UnionId).First(&loginUserInfo).Error
	if err != nil {
		//	走注册逻辑或者失败逻辑
		user := system.SysUser{
			Username:    FsUserInfo.TenantKey,
			Password:    "123456",
			NickName:    FsUserInfo.Name,
			AuthorityId: global.GlobalConfig.AuthorityID,
			Authorities: []system.SysAuthority{
				{
					AuthorityId: global.GlobalConfig.AuthorityID,
				},
			},
			Projects: []*system.Project{
				{
					GVA_MODEL: global2.GVA_MODEL{
						ID: global.GlobalConfig.ProjectID,
					},
				},
			},
		}
		gvaInfo, err := userService.Register(user)

		if err != nil {
			//	走注册失败逻辑
			return gvaInfo, err
		}
		FsUserInfo.GvaUserId = gvaInfo.ID
		err = global2.GVA_DB.Create(&FsUserInfo).Error
		if err != nil {
			//	走创建失败逻辑
			return gvaInfo, err
		}
		return gvaInfo, err
	} else {
		//走登录逻辑
		err := global2.GVA_DB.First(&gvaInfo, "id = ?", loginUserInfo.GvaUserId).Error
		if err != nil {
			//	登录失败
			return gvaInfo, err
		}
		return gvaInfo, err
	}
}

func (e *FsLoginService) BindFs(FsUserInfo model.FsUserInfo, id uint) (err error) {
	FsUserInfo.GvaUserId = id
	err = global2.GVA_DB.First(&FsUserInfo, "gva_user_id = ?", id).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		err = global2.GVA_DB.First(&FsUserInfo, "union_id = ?", FsUserInfo.UnionId).Error
		if err == nil {
			return errors.New("此飞书已绑定其他账号")
		}
		err = global2.GVA_DB.Create(&FsUserInfo).Error
		return err
	} else {
		return errors.New("请勿重复绑定")
	}
}
