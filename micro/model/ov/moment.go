package ov

import "time"

type Moment struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `gorm:"type:varchar(500)" json:"content"`
	ImageUrl  string    `json:"image_url"` //图片
	Mood      Mood      `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName  string    `gorm:"type:varchar(20)" json:"mood_name"`
	Tags      []Tag     `gorm:"many2many:moment_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User      User      `json:"user"`
	UserID    uint64    `json:"user_id"`
	ActionCount
	Permission uint8 `gorm:"type:smallint;default:0" json:"permission"` //查看权限
	Status     uint8 `gorm:"type:smallint;default:0" json:"-"`
}
