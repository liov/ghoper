package model

import "time"

type Dictionary struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Type      string     `gorm:"type:varchar(128)" json:"type"`
	Key       string     `gorm:"type:varchar(128)" json:"key"`
	Value     string     `gorm:"type:varchar(256)" json:"value"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	UpdatedAt *time.Time `json:"updated_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type Logger struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	RefID     uint64     `json:"ref_id"`
	RefKind   string     `gorm:"type:varchar(128)" json:"ref_kind"`
	Operation string     `gorm:"type:varchar(128)" json:"key"`
	Sequence  uint8      `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
	UpdatedAt *time.Time `json:"updated_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}
