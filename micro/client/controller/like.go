package controller

import (
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"time"
)

type Favorites struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `gorm:"type:varchar(20)" json:"name"`
	User      *User      `json:"user"`
	UserID    uint       `json:"user_id"`
	Count     uint       `json:"count"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type FavoritesCollection struct {
	Collection    Collection `json:"collection"`
	FavoritesName string     `json:"favorites"`
}

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	RefID     uint       `json:"ref_id"`
	Kind      string     `gorm:"type:varchar(10)" json:"kind"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

type Like struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	User      *User      `json:"user"`
	UserID    uint       `json:"user_id"`
	Count     uint       `json:"count"`
	Articles  []Article  `gorm:"many2many:article_like" json:"articles"`
	Moments   []Moment   `gorm:"many2many:moment_like" json:"moments"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

func AddLike(ctx iris.Context) {

}
func DelLike(ctx iris.Context) {

}

func AddCollection(ctx iris.Context) {

}

func DelCollection(ctx iris.Context) {

}

func GetCollection(ctx iris.Context) {
	id := ctx.Values().Get("userId").(uint)
	var favorites []Favorites
	initialize.DB.Where("user_id=?", id).Find(&favorites)
	common.Response(ctx, favorites)
}
