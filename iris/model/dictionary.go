package model

import "time"

type Dictionary struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `gorm:"type:varchar(128)" json:"type"`
	Key       string    `gorm:"type:varchar(128)" json:"key"`
	Value     string    `gorm:"type:varchar(256)" json:"value"`
	Status    uint8     `gorm:"type:smallint;default:0" json:"status"`
}
