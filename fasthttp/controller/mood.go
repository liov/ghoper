package controller

import (
	"service/initialize"
	"time"
)

type Mood struct {
	Name          string    `gorm:"type:varchar(20);primary_key" json:"name"`
	Description   string    `gorm:"type:varchar(100)" json:"description"`
	ExpressionURL string    `gorm:"type:varchar(100)" json:"expression_url"`
	Count         uint      `gorm:"default:0" json:"count"`
	CreatedAt     time.Time `json:"created_at"`
}

func ExistMoodByName(name string) *Mood {
	var mood Mood
	initialize.DB.Select("name,count").Where("name = ?", name).First(&mood)
	if mood.Name != "" {
		return &mood
	}

	return nil
}
