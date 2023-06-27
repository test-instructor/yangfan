package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"`              // 主键ID
	CreatedAt time.Time      `json:"CreatedAt,omitempty"`     // 创建时间
	UpdatedAt time.Time      `json:"UpdatedAt,omitempty"`     // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" json:"-"` // 删除时间
}
