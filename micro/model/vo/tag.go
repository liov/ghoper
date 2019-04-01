package vo

type Tag struct {
	Name        string `gorm:"type:varchar(10);primary_key" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
}

type CreatUser struct {
	CreatedBy User `json:"created_by"`
}

type ExtraInfo struct {
	UserID   uint64 `json:"user_id"`
	Sequence uint8  `gorm:"type:smallint;default:0" json:"sequence"` //排序，置顶
}

type Mood struct {
	Name          string `gorm:"type:varchar(20);primary_key" json:"name"`
	Description   string `gorm:"type:varchar(100)" json:"description"`
	ExpressionURL string `gorm:"type:varchar(100)" json:"expression_url"`
}

type Category struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`                               //直接父分类的ID
	Sequence uint8  `gorm:"type:smallint;default:0" json:"sequence"` //同级别的分类可根据sequence的值来排序，置顶
}
