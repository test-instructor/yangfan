package platform

import (
	"context"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	platformReq "github.com/test-instructor/yangfan/server/v2/model/platform/request"
)

type RunnerNodeService struct{}

// CreateRunnerNode 创建节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) CreateRunnerNode(ctx context.Context, rn *platform.RunnerNode) (err error) {
	err = global.GVA_DB.Create(rn).Error
	return err
}

// DeleteRunnerNode 删除节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) DeleteRunnerNode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&platform.RunnerNode{}, "id = ?", ID).Error
	return err
}

// DeleteRunnerNodeByIds 批量删除节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) DeleteRunnerNodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]platform.RunnerNode{}, "id in ?", IDs).Error
	return err
}

// UpdateRunnerNode 更新节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) UpdateRunnerNode(ctx context.Context, rn platform.RunnerNode) (err error) {
	err = global.GVA_DB.Model(&platform.RunnerNode{}).Where("id = ?", rn.ID).Updates(&rn).Error
	return err
}

// GetRunnerNode 根据ID获取节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) GetRunnerNode(ctx context.Context, ID string) (rn platform.RunnerNode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&rn).Error
	return
}

// GetRunnerNodeInfoList 分页获取节点记录
// Author [yourname](https://github.com/yourname)
func (rnService *RunnerNodeService) GetRunnerNodeInfoList(ctx context.Context, info platformReq.RunnerNodeSearch) (list []platform.RunnerNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&platform.RunnerNode{})
	var rns []platform.RunnerNode
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if len(info.RunContents) > 0 {
		db = db.Where("run_content IN ?", info.RunContents)
	} else if info.RunContent != "" {
		db = db.Where("run_content = ?", info.RunContent)
	}
	db.Order("id desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&rns).Error
	return rns, total, err
}
func (rnService *RunnerNodeService) GetRunnerNodePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
