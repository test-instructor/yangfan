package initialize

import (
	"fmt"
	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// migrateDCMCountColumns 将 lc_data_category_management 的 count/available_count 列变更为 JSON 类型（幂等）
func migrateDCMCountColumns(db *gorm.DB) error {
	dial := db.Dialector.Name()
	switch dial {
	case "mysql":
		// MySQL: 将列修改为 JSON 类型（已是 JSON 时不会改变）；忽略错误以保证幂等
		if err := db.Exec("ALTER TABLE `lc_data_category_management` MODIFY COLUMN `count` JSON NULL").Error; err != nil {
			global.GVA_LOG.Debug("migrate dcm count column (mysql) skipped", zap.Any("err", err))
		}
		if err := db.Exec("ALTER TABLE `lc_data_category_management` MODIFY COLUMN `available_count` JSON NULL").Error; err != nil {
			global.GVA_LOG.Debug("migrate dcm available_count column (mysql) skipped", zap.Any("err", err))
		}
	case "postgres":
		// Postgres: 使用 jsonb，并尝试把旧值转换为对象或空对象
		if err := db.Exec(`
			ALTER TABLE lc_data_category_management
			ALTER COLUMN count TYPE jsonb USING 
				CASE 
					WHEN jsonb_typeof(count::jsonb) IS NOT NULL THEN count::jsonb
					ELSE to_jsonb(count)
				END,
			ALTER COLUMN available_count TYPE jsonb USING 
				CASE 
					WHEN jsonb_typeof(available_count::jsonb) IS NOT NULL THEN available_count::jsonb
					ELSE to_jsonb(available_count)
				END;
		`).Error; err != nil {
			global.GVA_LOG.Debug("migrate dcm columns (postgres) skipped", zap.Any("err", err))
		}
	default:
		global.GVA_LOG.Info(fmt.Sprintf("skip dcm columns migration for dialector: %s", dial))
	}
	return nil
}
