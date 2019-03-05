package model

import "time"

// School 学校(教育经历)
type School struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Name       string     `gorm:"type:varchar(20)" json:"name"`
	Speciality string     `gorm:"type:varchar(100)" json:"speciality"` //专业
	StartTime  time.Time  `json:"start_time"`
	EndTime    time.Time  `json:"end_time"`
	User       *User      `json:"user"`
	UserID     uint       `json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	Status     uint8      `gorm:"type:smallint;default:0" json:"status"`
}

// MaxSchoolNameLen 学校或教育机构名的最大长度
const MaxSchoolNameLen = 200

// MaxSchoolSpecialityLen 专业的最大长度
const MaxSchoolSpecialityLen = 200
