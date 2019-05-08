package ov

type Tag struct {
	Name        string `gorm:"type:varchar(10);primary_key" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Status      uint8  `gorm:"type:smallint;default:0" json:"-"`
}

type Mood struct {
	Name          string `gorm:"type:varchar(20);primary_key" json:"name"`
	Description   string `gorm:"type:varchar(100)" json:"description"`
	ExpressionURL string `gorm:"type:varchar(100)" json:"expression_url"`
	Status        uint8  `gorm:"type:smallint;default:0" json:"-"`
}

type Category struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	Name     string `gorm:"unique_index" json:"name"`
	ParentID int    `json:"parent_id"`                               //直接父分类的ID
	Sequence uint8  `gorm:"type:smallint;default:0" json:"sequence"` //同级别的分类可根据sequence的值来排序，置顶
	Status   uint8  `gorm:"type:smallint;default:0" json:"-"`
}
