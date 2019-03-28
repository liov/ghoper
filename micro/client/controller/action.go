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
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    uint8      `gorm:"type:smallint;default:0" json:"status"`
}

//私有都不会传过去，why？iris传的时候用反射？
type UserAction struct {
	Collect []int64 `json:"collect"`
	Like    []int64 `json:"like"`
	Approve []int64 `json:"approve"`
	Comment []int64 `json:"comment"`
}

var IndexKind = map[int8]string{
	kindMoment:    "Moment",
	kindArticle:   "Article",
	kindDiary:     "Diary",
	kindDiaryBook: "DiaryBook",
}

var KindIndex = map[string]int8{
	"Moment":    kindMoment,
	"Article":   kindArticle,
	"Diary":     kindDiary,
	"DiaryBook": kindDiaryBook,
}

const (
	kindMoment = iota + 1
	kindArticle
	kindDiary
	kindDiaryBook
)

const (
	actionCollect = iota
	actionLike
	actionApprove
	actionBrowse
	actionComment
)

var IndexAction = map[int8]string{
	actionCollect: "Collect",
	actionLike:    "Like",
	actionApprove: "Approve",
	actionBrowse:  "Browse",
	actionComment: "Comment",
}

func setCountToRedis(userId uint64, refId uint64, kind string, action int8, num int8) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("SELECT", KindIndex[kind])
	if num > 0 {
		conn.Send("SADD", strings.Join([]string{"User", strconv.FormatUint(userId, 10), kind, IndexAction[action]}, "_"), refId)
	} else {
		conn.Send("SREM", strings.Join([]string{"User", strconv.FormatUint(userId, 10), kind, IndexAction[action]}, "_"), refId)
	}
	conn.Send("HINCRBY", strings.Join([]string{kind, strconv.FormatUint(refId, 10), "Action", "Count"}, "_"), IndexAction[action], num)
	conn.Send("SELECT", 0)
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error("缓存失败", err)
	}
	return nil
}

func getRedisAction(userId string, kind string) *UserAction {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := strings.Join([]string{"User", userId, kind}, "_")
	conn.Send("SELECT", KindIndex[kind])
	conn.Send("SMEMBERS", key+"_Collect")
	conn.Send("SMEMBERS", key+"Comment")
	conn.Send("SMEMBERS", key+"_Like")
	conn.Send("SMEMBERS", key+"_Approve")
	conn.Send("SMEMBERS", key+"_Browse")
	conn.Send("SELECT", 0)
	conn.Flush()
	conn.Receive()
	userAction := new(UserAction)
	collect, err := redis.Int64s(conn.Receive())
	userAction.Collect = collect
	comment, err := redis.Int64s(conn.Receive())
	userAction.Collect = comment
	like, err := redis.Int64s(conn.Receive())
	userAction.Like = like
	approve, err := redis.Int64s(conn.Receive())
	userAction.Approve = approve
	browse, err := redis.Int64s(conn.Receive())
	userAction.Approve = browse
	conn.Receive()
	if err != nil {
		golog.Error(err)
	}
	return userAction
}

type ActionCount struct {
	CollectCount int64 `gorm:"default:0" json:"collect_count"` //收藏
	LikeCount    int64 `gorm:"default:0" json:"like_count"`    //喜欢
	ApproveCount int64 `gorm:"default:0" json:"approve_count"` //点赞
	CommentCount int64 `gorm:"default:0" json:"comment_count"` //评论
	BrowseCount  int64 `gorm:"default:0" json:"browse_count"`  //浏览
}

func getActionCount(refId uint64, kind int8) *ActionCount {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	id := strconv.FormatUint(refId, 10)
	key := strings.Join([]string{IndexKind[kind], id, "Action", "Count"}, "_")
	conn.Send("SELECT", kind)
	conn.Send("HGETALL", key)
	conn.Send("SELECT", 0)
	conn.Flush()
	conn.Receive()
	actionCount := new(ActionCount)
	action, err := redis.Int64Map(conn.Receive())
	actionCount.CollectCount = action["Collect"]
	actionCount.LikeCount = action["Like"]
	actionCount.ApproveCount = action["Approve"]
	actionCount.BrowseCount = action["Browse"]
	conn.Receive()
	if err != nil {
		golog.Error(err)
	}
	return actionCount
}

//数据量大，每个用户维护一张喜欢表
func AddLike(ctx iris.Context) {
	var like Like
	if err := ctx.ReadJSON(&like); err != nil {
		common.Response(ctx, "参数无效")
		return
	}
	userId := ctx.Values().Get("userId").(uint64)
	var count int
	initialize.DB.Model(&model.Like{}).Where("ref_id =? AND kind = ?", like.RefID, like.Kind).Count(&count)
	if count > 0 {
		initialize.DB.Delete(&like)
		setCountToRedis(userId, like.RefID, like.Kind, actionLike, -1)
		common.Response(ctx, "成功", e.SUCCESS)
		return
	}

	like.UserID = userId
	like.Status = 1
	err := initialize.DB.Create(&like).Error
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "喜欢失败", e.ERROR)
		return
	}

	setCountToRedis(userId, like.RefID, like.Kind, actionLike, 1)

	common.Response(ctx, "成功", e.SUCCESS)
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

	setCountToRedis(userId, fc.RefID, fc.Kind, actionCollect, 1)

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
