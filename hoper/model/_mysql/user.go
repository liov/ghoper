package _mysql

import (
	"crypto/md5"
	"fastService/initialize"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	ActivatedAt     *time.Time `json:"activated_at"` //激活时间
	Name            string     `gorm:"type:varchar(10);not null" json:"name"`
	Password        string     `json:"password"`
	Email           string     `gorm:"type:varchar(20);unique_index;not null" json:"email"`
	Phone           *string    `gorm:"type:varchar(20);unique_index" json:"phone"` //手机号
	Sex             uint8      `gorm:"type:tinyint(1) unsigned;not null" json:"sex"`
	Introduce       string     `gorm:"type:varchar(500)" json:"introduce"`             //简介
	Score           uint       `gorm:default:0" json:"score"`                          //积分
	Signature       string     `gorm:"type:varchar(100)" json:"signature"`             //个人签名
	Role            uint8      `gorm:"type:tinyint(1) unsigned;default:0" json:"role"` //管理员or用户
	AvatarURL       string     `gorm:"type:varchar(100)" json:"avatar_url"`            //头像
	CoverURL        string     `gorm:"type:varchar(100)" json:"cover_url"`             //个人主页背景图片URL
	Address         string     `gorm:"type:varchar(100)" json:"address"`
	Location        string     `gorm:"type:varchar(100)" json:"location"`
	Schools         []School   `json:"schools"` //教育经历
	Careers         []Career   `json:"careers"` //职业经历
	UpdatedAt       *time.Time `json:"updated_at"`
	BannedAt        *time.Time `sql:"index" json:"banned_at"`
	CreatedAt       time.Time  `json:"created_at"`
	LastActivatedAt *time.Time `json:"last_activated_at"`                                //激活时间
	LastName        string     `gorm:"type:varchar(100)" json:"last_name"`               //上个名字
	Status          uint8      `gorm:"type:tinyint(1) unsigned;default:0" json:"status"` //状态
	Like            Love       `json:"love"`                                             //和Collection挺像的，不过一个User可以对应多个C，只能对应一个L
	LikeID          uint       `json:"love_id"`
	/*	Follows        []*User      `gorm:"-" json:"follows"`                //gorm:"foreignkey:FollowID []Follow里的User
		Followeds      []*User      `gorm:"-" json:"followeds"`              //gorm:"foreignkey:UserID"	[]Follow里的FollowUser*/
	FollowCount    uint         `gorm:"default:0" json:"follow_count"`   //关注数量
	FollowedCount  uint         `gorm:"default:0" json:"followed_count"` //被关注数量
	ArticleCount   uint         `gorm:"default:0" json:"article_count"`  //文章数量
	MomentCount    uint         `gorm:"default:0" json:"moment_count"`
	DiaryBookCount uint         `gorm:"default:0" json:"diary_book_count"`
	DiaryCount     uint         `gorm:"default:0" json:"diary_count"`
	CommentCount   uint         `gorm:"default:0" json:"comment_count"`               //评论数量
	Collections    []Collection `gorm:"many2many:user_collection" json:"collections"` //收藏夹？
	Articles       []Article    `json:"articles"`
	Moments        []Moment     `json:"moments"`
	DiaryBooks     []DiaryBook  `json:"diary_books"`
	Diarys         []Diary      `json:"diarys"`
}

type Follow struct {
	User         User       `json:"user"`
	UserID       uint       `gorm:"primary_key" json:"user_id"` //一个关注另一个，ID小的做UserID
	FollowUser   User       `json:"follow_user"`
	FollowID     uint       `gorm:"primary_key" json:"follow_id"`
	FollowUserAt *time.Time //FollowUser关注User时间
	UserFollowAt *time.Time //User关注FollowUser时间                                                 //互相关注时间
	Status       uint8      `gorm:"type:tinyint(1) unsigned;default:0" json:"status"` //0 都生效，1前面生效，2后面生效，3都不生效
}

// CheckPassword 验证密码是否正确
func (user User) CheckPassword(password string) bool {
	if password == "" || user.Password == "" {
		return false
	}
	return user.EncryptPassword(password, user.Salt()) == user.Password
}

// Salt 每个用户都有一个不同的盐
func (user User) Salt() string {
	var userSalt string
	if user.Password == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Password[0:5]
	}
	return userSalt
}

// EncryptPassword 给密码加密
func (user User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + initialize.ServerSettings.PassSalt
	hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return
}

const (
	// UserRoleNormal 普通用户
	UserRoleNormal = iota

	// UserRoleEditor 网站编辑
	UserRoleEditor

	// UserRoleAdmin 管理员
	UserRoleAdmin

	// UserRoleSuperAdmin 超级管理员
	UserRoleSuperAdmin

	// UserRoleCrawler 爬虫，网站编辑或管理员登陆后台后，操作爬虫去抓取文章
	// 这时，生成的文章，其作者是爬虫账号。没有直接使用爬虫账号去登陆的情况.
	UserRoleCrawler
)

const (
	// UserStatusInActive 未激活
	UserStatusInActive = iota

	// UserStatusActived 已激活
	UserStatusActived

	// UserStatusFrozen 已冻结
	UserStatusFrozen
)

const (
	// UserSexMale 男
	UserSexMale = iota

	// UserSexFemale 女
	UserSexFemale
	//未填写
	UserSexNil
)
const (
	// MaxUserNameLen 用户名的最大长度
	MaxUserNameLen = 20

	// MinUserNameLen 用户名的最小长度
	MinUserNameLen = 4

	// MaxPassLen 密码的最大长度
	MaxPassLen = 20

	// MinPassLen 密码的最小长度
	MinPassLen = 6

	// MaxSignatureLen 个性签名最大长度
	MaxSignatureLen = 200

	// MaxLocationLen 居住地的最大长度
	MaxLocationLen = 200

	// MaxIntroduceLen 个人简介的最大长度
	MaxIntroduceLen = 500
)

const (
	// ActiveTime 生成激活账号的链接
	ActiveTime = "activeTime"

	// ResetTime 生成重置密码的链接
	ResetTime = "resetTime"

	// LoginUser 用户信息
	LoginUser = "loginUser"

	// ArticleMinuteLimit 用户每分钟最多能发表的文章数
	MomentMinuteLimit = "momentMinuteLimit"

	// ArticleDayLimit 用户每天最多能发表的文章数
	MomentDayLimit = "momentDayLimit"

	// ArticleMinuteLimit 用户每分钟最多能发表的文章数
	ArticleMinuteLimit = "articleMinuteLimit"

	// ArticleDayLimit 用户每天最多能发表的文章数
	ArticleDayLimit = "articleDayLimit"

	// CommentMinuteLimit 用户每分钟最多能发表的评论数
	CommentMinuteLimit = "commentMinuteLimit"

	// CommentDayLimit 用户每天最多能发表的评论数
	CommentDayLimit = "commentDayLimit"
)

// 积分相关常量
const (
	// ArticleScore 创建话题时增加的积分
	ArticleScore = 5

	// ByCommentScore 话题或投票被评论时增加的积分
	ByCommentScore = 2

	// ByCollectScore 话题或投票被收藏时增加的积分
	ByCollectScore = 2

	// CommentScore 评论话题或投票时增加的积分
	CommentScore = 1

	// CollectScore 收藏话题或投票时增加的积分
	CollectScore = 1
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
