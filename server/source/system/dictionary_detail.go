package system

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	sysModel "github.com/test-instructor/yangfan/server/v2/model/system"
	"github.com/test-instructor/yangfan/server/v2/service/system"
	"gorm.io/gorm"
)

const initOrderDictDetail = initOrderDict + 1

type initDictDetail struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDictDetail, &initDictDetail{})
}

func (i *initDictDetail) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysDictionaryDetail{})
}

func (i *initDictDetail) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysDictionaryDetail{})
}

func (i *initDictDetail) InitializerName() string {
	return sysModel.SysDictionaryDetail{}.TableName()
}

func (i *initDictDetail) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	dicts, ok := ctx.Value(new(initDict).InitializerName()).([]sysModel.SysDictionary)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext,
			fmt.Sprintf("未找到 %s 表初始化数据", sysModel.SysDictionary{}.TableName()))
	}
	True := true
	dicts[0].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "男", Value: "1", Status: &True, Sort: 1},
		{Label: "女", Value: "2", Status: &True, Sort: 2},
	}

	dicts[1].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "smallint", Value: "1", Status: &True, Extend: "mysql", Sort: 1},
		{Label: "mediumint", Value: "2", Status: &True, Extend: "mysql", Sort: 2},
		{Label: "int", Value: "3", Status: &True, Extend: "mysql", Sort: 3},
		{Label: "bigint", Value: "4", Status: &True, Extend: "mysql", Sort: 4},
		{Label: "int2", Value: "5", Status: &True, Extend: "pgsql", Sort: 5},
		{Label: "int4", Value: "6", Status: &True, Extend: "pgsql", Sort: 6},
		{Label: "int6", Value: "7", Status: &True, Extend: "pgsql", Sort: 7},
		{Label: "int8", Value: "8", Status: &True, Extend: "pgsql", Sort: 8},
	}

	dicts[2].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "date", Value: "0", Status: &True, Extend: "mysql", Sort: 0},
		{Label: "time", Value: "1", Status: &True, Extend: "mysql", Sort: 1},
		{Label: "year", Value: "2", Status: &True, Extend: "mysql", Sort: 2},
		{Label: "datetime", Value: "3", Status: &True, Extend: "mysql", Sort: 3},
		{Label: "timestamp", Value: "5", Status: &True, Extend: "mysql", Sort: 5},
		{Label: "timestamptz", Value: "6", Status: &True, Extend: "pgsql", Sort: 5},
	}
	dicts[3].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "float", Value: "0", Status: &True, Extend: "mysql", Sort: 0},
		{Label: "double", Value: "1", Status: &True, Extend: "mysql", Sort: 1},
		{Label: "decimal", Value: "2", Status: &True, Extend: "mysql", Sort: 2},
		{Label: "numeric", Value: "3", Status: &True, Extend: "pgsql", Sort: 3},
		{Label: "smallserial", Value: "4", Status: &True, Extend: "pgsql", Sort: 4},
	}

	dicts[4].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "char", Value: "0", Status: &True, Extend: "mysql", Sort: 0},
		{Label: "varchar", Value: "1", Status: &True, Extend: "mysql", Sort: 1},
		{Label: "tinyblob", Value: "2", Status: &True, Extend: "mysql", Sort: 2},
		{Label: "tinytext", Value: "3", Status: &True, Extend: "mysql", Sort: 3},
		{Label: "text", Value: "4", Status: &True, Extend: "mysql", Sort: 4},
		{Label: "blob", Value: "5", Status: &True, Extend: "mysql", Sort: 5},
		{Label: "mediumblob", Value: "6", Status: &True, Extend: "mysql", Sort: 6},
		{Label: "mediumtext", Value: "7", Status: &True, Extend: "mysql", Sort: 7},
		{Label: "longblob", Value: "8", Status: &True, Extend: "mysql", Sort: 8},
		{Label: "longtext", Value: "9", Status: &True, Extend: "mysql", Sort: 9},
	}

	dicts[5].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "tinyint", Value: "1", Extend: "mysql", Status: &True},
		{Label: "bool", Value: "2", Extend: "pgsql", Status: &True},
	}

	dicts[6].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "Host", Value: "Host", Status: &True, Sort: 1},
		{Label: "Token", Value: "Token", Status: &True, Sort: 2},
	}

	dicts[8].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "String", Value: "String", Status: &True, Sort: 1},
		{Label: "Integer", Value: "Integer", Status: &True, Sort: 2},
		{Label: "Float", Value: "Float", Status: &True, Sort: 3},
		{Label: "Boolean", Value: "Boolean", Status: &True, Sort: 4},
		{Label: "List", Value: "List", Status: &True, Sort: 5},
		{Label: "Dict", Value: "Dict", Status: &True, Sort: 6},
	}

	dicts[9].SysDictionaryDetails = []sysModel.SysDictionaryDetail{
		{Label: "等于", Value: "equals", Status: &True, Sort: 1},
		{Label: "大于", Value: "greater_than", Status: &True, Sort: 2},
		{Label: "小于", Value: "less_than", Status: &True, Sort: 3},
		{Label: "大于等于", Value: "greater_or_equals", Status: &True, Sort: 4},
		{Label: "小于等于", Value: "less_or_equals", Status: &True, Sort: 5},
		{Label: "不等于", Value: "not_equal", Status: &True, Sort: 6},
		{Label: "包含", Value: "contains", Status: &True, Sort: 7},
		{Label: "类型匹配", Value: "type_match", Status: &True, Sort: 8},
		{Label: "正则匹配", Value: "regex_match", Status: &True, Sort: 9},
		{Label: "长度等于", Value: "length_equals", Status: &True, Sort: 10},
		{Label: "被包含", Value: "contained_by", Status: &True, Sort: 11},
		{Label: "长度小于", Value: "length_less_than", Status: &True, Sort: 12},
		{Label: "字符串相等", Value: "string_equals", Status: &True, Sort: 13},
		{Label: "忽略大小写相等", Value: "equal_fold", Status: &True, Sort: 14},
		{Label: "长度小于等于", Value: "length_less_or_equals", Status: &True, Sort: 15},
		{Label: "长度大于", Value: "length_greater_than", Status: &True, Sort: 16},
		{Label: "长度大于等于", Value: "length_greater_or_equals", Status: &True, Sort: 17},
		{Label: "开头字符串", Value: "startswith", Status: &True, Sort: 18},
		{Label: "结尾字符串", Value: "endswith", Status: &True, Sort: 19},
	}
	for _, dict := range dicts {
		if err := db.Model(&dict).Association("SysDictionaryDetails").
			Replace(dict.SysDictionaryDetails); err != nil {
			return ctx, errors.Wrap(err, sysModel.SysDictionaryDetail{}.TableName()+"表数据初始化失败!")
		}
	}
	return ctx, nil
}

func (i *initDictDetail) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var dict sysModel.SysDictionary
	if err := db.Preload("SysDictionaryDetails").
		First(&dict, &sysModel.SysDictionary{Name: "数据库bool类型"}).Error; err != nil {
		return false
	}
	return len(dict.SysDictionaryDetails) > 0 && dict.SysDictionaryDetails[0].Label == "tinyint"
}
