package vo

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(10);not null" json:"name"`
	Sex       string `gorm:"type:varchar(1);not null" json:"sex"`
	Signature string `gorm:"type:varchar(100)" json:"signature"`  //个人签名
	AvatarURL string `gorm:"type:varchar(100)" json:"avatar_url"` //头像
}

type UserData struct {
	Email        string     `gorm:"type:varchar(20);unique_index;not null" json:"email"`
	Phone        *string    `gorm:"type:varchar(20);unique_index" json:"phone"` //手机号
	Birthday     *time.Time `json:"birthday"`
	Introduction string     `gorm:"type:varchar(500)" json:"introduction"` //简介
	Score        uint64     `gorm:"default:0" json:"score"`                //积分
}

type UserMore struct {
	Address  string      `gorm:"type:varchar(100)" json:"address"`
	Location string      `gorm:"type:varchar(100)" json:"location"`
	EduExps  []Education `json:"edu_exps"`  //教育经历
	WorkExps []Work      `json:"work_exps"` //职业经历
}

type UserNoPub struct {
	ActivatedAt     *time.Time `json:"activated_at"`                        //激活时间
	Role            uint8      `gorm:"type:smallint;default:0" json:"role"` //管理员or用户
	BannedAt        *time.Time `sql:"index" json:"banned_at"`
	LastActivatedAt *time.Time `json:"last_activated_at"`                  //最后活跃时间
	LastName        string     `gorm:"type:varchar(100)" json:"last_name"` //上个名字
}

type UserOwn struct {
	//和Collection挺像的，不过一个User可以对应多个C，只能对应一个L
	//一个Like似乎没用啊，一个人的喜欢可以是多个，收藏也是这样,如果说分表，一个喜欢夹对应多条喜欢，一个索引作用没什么意义，为什么不存一个表里
	Collections []Collection `json:"collections"`
	Follows     []*User      `gorm:"-" json:"follows"`   //gorm:"foreignkey:FollowID []Follow里的User
	Followeds   []*User      `gorm:"-" json:"followeds"` //gorm:"foreignkey:UserID"	[]Follow里的FollowUser
	Favorites   []Favorites  `json:"favorites"`          //收藏夹？
}

type UserOwnCount struct {
	FollowCount    uint64 `gorm:"default:0" json:"follow_count"`   //关注数量
	FollowedCount  uint64 `gorm:"default:0" json:"followed_count"` //被关注数量
	ArticleCount   uint64 `gorm:"default:0" json:"article_count"`  //文章数量
	MomentCount    uint64 `gorm:"default:0" json:"moment_count"`
	DiaryBookCount uint64 `gorm:"default:0" json:"diary_book_count"`
	DiaryCount     uint64 `gorm:"default:0" json:"diary_count"`
	CommentCount   uint64 `gorm:"default:0" json:"comment_count"` //评论数量
}

type Follow struct {
	UserID   uint64 `gorm:"primary_key" json:"user_id"` //一个关注另一个，ID小的做UserID
	FollowID uint64 `gorm:"primary_key" json:"follow_id"`
}

type FollowNoPub struct {
	User       User `json:"user"`
	FollowUser User `json:"follow_user"`
}

type FollowStatus struct {
	FollowUserAt *time.Time //FollowUser关注User时间
	UserFollowAt *time.Time //User关注FollowUser时间                                                //互相关注时间
	Status       uint8      `gorm:"type:smallint;default:0" json:"status"` //0 都生效，1前面生效，2后面生效，3都不生效
}
