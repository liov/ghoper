package controller

import "time"

// Education 教育经历
type Education struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	School     string     `gorm:"type:varchar(20)" json:"school"`
	Speciality string     `gorm:"type:varchar(100)" json:"speciality"` //专业
	StartTime  time.Time  `json:"start_time"`
	EndTime    time.Time  `json:"end_time"`
	User       *User      `json:"user"`
	UserID     uint64     `json:"user_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	Status     uint8      `gorm:"type:smallint;default:0" json:"status"`
}

// Work 职业生涯
type Work struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Company   string     `json:"company"` //公司或组织
	Title     string     `json:"title"`   //职位
	UserID    uint64     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

// MaxCareerCompanyLen 公司或组织名称的最大长度
const MaxCareerCompanyLen = 200

// MaxCareerTitleLen 职位的最大长度
const MaxCareerTitleLen = 200

type School struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(20)" json:"name"`
	CreatedTime time.Time  `json:"start_time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	Status      uint8      `gorm:"type:smallint;default:0" json:"status"`
}

// MaxSchoolNameLen 学校或教育机构名的最大长度
const MaxSchoolNameLen = 200

// MaxSchoolSpecialityLen 专业的最大长度
const MaxSchoolSpecialityLen = 200
