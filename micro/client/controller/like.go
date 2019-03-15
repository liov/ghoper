package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/initialize"
	"hoper/model"
	"strconv"
	"time"
)

type Favorites struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	Name        string     `gorm:"type:varchar(20)" json:"name"`
	User        User       `json:"user"`
	UserID      uint       `json:"user_id"`
	FollowUsers []User     `json:"follow_users"`
	Count       uint       `json:"count"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	Status      uint8      `gorm:"type:smallint;default:0" json:"status"`
}

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID          uint        `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	RefID       uint        `json:"ref_id"`
	Kind        string      `gorm:"type:varchar(10)" json:"kind"`
	Favorites   []Favorites `json:"favorites"`
	FavoritesID uint        `json:"favorites_id"`
	UserID      uint        `json:"user_id"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `sql:"index" json:"deleted_at"`
	Status      uint8       `gorm:"type:smallint;default:0" json:"status"`
}

type Like struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	RefID     uint       `json:"ref_id"`
	Kind      string     `gorm:"type:varchar(10)" json:"kind"`
	User      User       `json:"user"`
	UserID    uint       `json:"user_id"`
	Count     uint       `json:"count"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

//数据量大，每个用户维护一张喜欢表
func AddLike(ctx iris.Context) {

}
func DelLike(ctx iris.Context) {

}

func AddCollection(ctx iris.Context) {
	type FavoritesCollection struct {
		RefID        uint   `json:"ref_id"`
		Kind         string `json:"kind"`
		FavoritesIDs []uint `json:"favorites_ids"`
	}

	var fc FavoritesCollection
	if err := ctx.ReadJSON(&fc); err != nil {
		common.Response(ctx, "参数无效")
		return
	}

	userId := ctx.Values().Get("userId").(int)
	var count int
	initialize.DB.Model(&model.Favorites{}).Where("user_id =? AND id in (?)", userId, fc.FavoritesIDs).Count(&count)
	if count != len(fc.FavoritesIDs) {
		common.Response(ctx, "收藏夹无效")
		return
	}

	initialize.DB.Model(&model.Collection{}).Where("ref_id =? AND kind = ? AND favorites_id in (?)", fc.RefID, fc.Kind, fc.FavoritesIDs).Count(&count)
	if count > 0 {
		common.Response(ctx, "已收藏")
		return
	}

	var err error
	for _, v := range fc.FavoritesIDs {
		err = initialize.DB.Create(&Collection{RefID: fc.RefID, Kind: fc.Kind, FavoritesID: v, Status: 1}).Error
	}
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "收藏失败", e.ERROR)
		return
	}

	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Do("SADD", "user"+strconv.Itoa(userId))
	common.Response(ctx, "收藏成功", e.SUCCESS)
}

func DelCollection(ctx iris.Context) {

}

func GetFavorite(ctx iris.Context) {
	id := ctx.Values().Get("userId").(uint)
	var favorites []Favorites
	initialize.DB.Where("user_id=?", id).Find(&favorites)
	common.Response(ctx, favorites)
}

func AddFavorite(ctx iris.Context) {
	userId := ctx.Values().Get("userId").(uint)

	var f Favorites
	if err := ctx.ReadJSON(&f); err != nil {
		common.Response(ctx, "参数无效")
		return
	}
	f.UserID = userId
	var count int
	initialize.DB.Model(&model.Favorites{}).Where(&f).Count(&count)
	if count > 0 {
		common.Response(ctx, "收藏夹已存在")
		return
	}
	initialize.DB.Create(&f)
	common.Response(ctx, f, "添加成功", e.SUCCESS)
}
