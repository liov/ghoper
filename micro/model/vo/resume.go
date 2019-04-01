package vo

import "time"

type Resume struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	Kind       uint8     `gorm:"type:smallint" json:"kind"`
	School     string    `gorm:"type:varchar(20)" json:"school"`
	Speciality string    `gorm:"type:varchar(100)" json:"speciality"` //专业
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	UserID     uint64    `json:"user_id"`
}

type ResumeOwn struct {
	Remark string `json:"remark"`
}

type Education struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	School     string    `gorm:"type:varchar(20)" json:"school"`
	Speciality string    `gorm:"type:varchar(100)" json:"speciality"` //专业
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	UserID     uint64    `json:"user_id"`
}

type Work struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	Company   string    `json:"company"` //公司或组织
	Title     string    `json:"title"`   //职位
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	UserID    uint64    `json:"user_id"`
}
