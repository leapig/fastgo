package entity

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

// TODO

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// Scan
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseEntity struct {
	ID        int64          `json:"id"           db:"id"        gorm:"column:id;primary_key;type:int"`
	CreatedAt time.Time      `json:"create_at"    db:"create_at" gorm:"column:create_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"update_at"    db:"update_at" gorm:"column:update_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"delete_at"    db:"delete_at" gorm:"column:delete_at;comment:删除时间"`
}
