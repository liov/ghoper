package ov

type User struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(10);not null" json:"name"`
	Sex       string `gorm:"type:varchar(1);not null" json:"sex"`
	Score     uint64 `gorm:"default:0" json:"score"`              //积分
	Signature string `gorm:"type:varchar(100)" json:"signature"`  //个人签名
	AvatarURL string `gorm:"type:varchar(100)" json:"avatar_url"` //头像
	Status    uint8  `gorm:"type:smallint;default:0" json:"-"`
}

type UserData struct {
	Account  string  `gorm:"type:varchar(20);unique_index" json:"account"`
	Email    string  `gorm:"type:varchar(20);unique_index;not null" json:"email"`
	Phone    *string `gorm:"type:varchar(20);unique_index" json:"phone"` //手机号
	Password string  `gorm:"type:varchar(100)" json:"-"`
}

type Follow struct {
	UserID   uint64 `gorm:"primary_key" json:"user_id"` //一个关注另一个，ID小的做UserID
	FollowID uint64 `gorm:"primary_key" json:"follow_id"`
}
