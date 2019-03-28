package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"strconv"
	"strings"
	"time"
)

type Favorites struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	Name        string     `gorm:"type:varchar(20)" json:"name"`
	User        User       `json:"user"`
	UserID      uint64     `json:"user_id"`
	FollowUsers []User     `json:"follow_users"`
	Count       uint64     `json:"count"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	Status      uint8      `gorm:"type:smallint;default:0" json:"status"`
}

//收藏夹？像网易云一样可以收藏别人的歌单
type Collection struct {
	ID          uint64      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	RefID       uint64      `json:"ref_id"`
	Kind        string      `gorm:"type:varchar(10)" json:"kind"`
	Favorites   []Favorites `json:"favorites"`
	FavoritesID uint64      `json:"favorites_id"`
	UserID      uint64      `json:"user_id"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `sql:"index" json:"deleted_at"`
	Status      uint8       `gorm:"type:smallint;default:0" json:"status"`
}

type Like struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	RefID     uint64     `json:"ref_id"`
	Kind      string     `gorm:"type:varchar(10)" json:"kind"`
	User      User       `json:"user"`
	UserID    uint64     `json:"user_id"`
	Count     uint64     `json:"count"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

//私有都不会传过去，why？iris传的时候用反射？
type UserLike struct {
	Collection []int64 `json:"collection"`
	Like       []int64 `json:"like"`
	Approve    []int64 `json:"approve"`
}

var kindIndex = map[string]int{
	"Moment":  1,
	"Article": 2,
	"Diary":   3,
}

func CountToRedis(userId uint64, refId uint64, kind string, operation string) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("SELECT", kindIndex[kind])
	conn.Send("SADD", strings.Join([]string{"User", strconv.FormatUint(userId, 10), kind, operation}, "_"), refId)
	conn.Send("INCR", strings.Join([]string{kind, strconv.FormatUint(refId, 10), operation, "Count"}, "_"))
	conn.Send("SELECT", 0)
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error("缓存失败", err)
	}
	return nil
}

func getRedisLike(userId string, kind string) *UserLike {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := strings.Join([]string{"User", userId, kind}, "_")
	conn.Send("SELECT", kindIndex[kind])
	conn.Send("SMEMBERS", key+"_Collect")
	conn.Send("SMEMBERS", key+"_Like")
	conn.Send("SMEMBERS", key+"_Approve")
	conn.Send("SELECT", 0)
	conn.Flush()
	conn.Receive()
	userLike := new(UserLike)
	collection, err := redis.Int64s(conn.Receive())
	userLike.Collection = collection
	like, err := redis.Int64s(conn.Receive())
	userLike.Like = like
	approve, err := redis.Int64s(conn.Receive())
	userLike.Approve = approve
	conn.Receive()
	if err != nil {
		golog.Error(err)
	}
	return userLike
}

type LikeCount struct {
	collection int64 `json:"collection"`
	like       int64 `json:"like"`
	approve    int64 `json:"approve"`
}

func getLikeCount(refId string, kind string) *LikeCount {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := strings.Join([]string{kind, refId}, "_")
	conn.Send("SELECT", kindIndex[kind])
	conn.Send("GET", key+"_Collect_Count")
	conn.Send("GET", key+"_Like_Count")
	conn.Send("GET", key+"_Approve_Count")
	conn.Send("SELECT", 0)
	conn.Flush()
	conn.Receive()
	likeCount := new(LikeCount)
	collection, err := redis.Int64(conn.Receive())
	likeCount.collection = collection
	like, err := redis.Int64(conn.Receive())
	likeCount.like = like
	approve, err := redis.Int64(conn.Receive())
	likeCount.approve = approve
	conn.Receive()
	if err != nil {
		golog.Error(err)
	}
	return likeCount
}

//数据量大，每个用户维护一张喜欢表
func AddLike(ctx iris.Context) {

}
func DelLike(ctx iris.Context) {

}

func AddCollection(ctx iris.Context) {
	type FavoritesCollection struct {
		RefID        uint64   `json:"ref_id"`
		Kind         string   `json:"kind"`
		FavoritesIDs []uint64 `json:"favorites_ids"`
	}

	var fc FavoritesCollection
	if err := ctx.ReadJSON(&fc); err != nil {
		common.Response(ctx, "参数无效")
		return
	}

	userId := ctx.Values().Get("userId").(uint64)
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
		err = initialize.DB.Create(&Collection{RefID: fc.RefID, Kind: fc.Kind, UserID: userId, FavoritesID: v, Status: 1}).Error
	}
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "收藏失败", e.ERROR)
		return
	}

	CountToRedis(userId, fc.RefID, fc.Kind, "Collect")

	common.Response(ctx, "收藏成功", e.SUCCESS)
}

func DelCollection(ctx iris.Context) {

}

func GetFavorite(ctx iris.Context) {
	id := ctx.Values().Get("userId").(uint64)
	var favorites []Favorites
	initialize.DB.Where("user_id=?", id).Find(&favorites)
	common.Response(ctx, favorites)
}

func AddFavorite(ctx iris.Context) {
	userId := ctx.Values().Get("userId").(uint64)

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
