package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/ov"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Action interface {
	GETUserID() uint64
	GETRefID() uint64
	GETKind() int8
	GETStatus() uint8
}

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

func (c *Collection) GETUserID() uint64 {
	return c.UserID
}

func (c *Collection) GETRefID() uint64 {
	return c.RefID
}

func (c *Collection) GETKind() int8 {
	return KindIndex[c.Kind]
}

func (c *Collection) GETStatus() uint8 {
	return c.Status
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

func (l *Like) GETUserID() uint64 {
	return l.UserID
}

func (l *Like) GETRefID() uint64 {
	return l.RefID
}

func (l *Like) GETKind() int8 {
	return KindIndex[l.Kind]
}

func (l *Like) GETStatus() uint8 {
	return l.Status
}

//私有都不会传过去，why？iris传的时候用反射？
type UserAction struct {
	Collect []int64 `json:"collect"`
	Like    []int64 `json:"like"`
	Approve []int64 `json:"approve"`
	Comment []int64 `json:"comment"`
	Browse  []int64 `json:"browse"`
}

var IndexKind = map[int8]string{
	kindMoment:    "Moment",
	kindArticle:   "Article",
	kindDiary:     "Diary",
	kindDiaryBook: "DiaryBook",
}

var KindIndex = map[string]int8{
	"moment":    kindMoment,
	"article":   kindArticle,
	"diary":     kindDiary,
	"diaryBook": kindDiaryBook,
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
	actionComment
	actionBrowse
)

var IndexAction = map[int8]string{
	actionCollect: "Collect",
	actionLike:    "Like",
	actionApprove: "Approve",
	actionComment: "Comment",
	actionBrowse:  "Browse",
}

func setCountToRedis(userID uint64, refId uint64, kind int8, action int8, num int8) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("SELECT", kind)

	if num > 0 {
		conn.Send("SADD", strings.Join([]string{"User", strconv.FormatUint(userID, 10), IndexKind[kind], IndexAction[action]}, "_"), refId)
	} else {
		conn.Send("SREM", strings.Join([]string{"User", strconv.FormatUint(userID, 10), IndexKind[kind], IndexAction[action]}, "_"), refId)
	}
	conn.Send("HINCRBY", strings.Join([]string{IndexKind[kind], strconv.FormatUint(refId, 10), "Action", "Count"}, "_"), IndexAction[action], num)
	conn.Send("ZINCRBY", strings.Join([]string{IndexKind[kind], strconv.FormatUint(refId, 10), IndexAction[action], "Sorted"}, "_"), num, refId)
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error("缓存失败:", err)
	}
	return nil
}

func getRedisAction(userID string, kind int8) *UserAction {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := strings.Join([]string{"User", userID, IndexKind[kind]}, "_")
	conn.Send("SELECT", kind)
	conn.Send("SMEMBERS", key+"_Collect")
	conn.Send("SMEMBERS", key+"Comment")
	conn.Send("SMEMBERS", key+"_Like")
	conn.Send("SMEMBERS", key+"_Approve")
	conn.Send("SMEMBERS", key+"_Browse")
	conn.Flush()
	conn.Receive()
	userAction := new(UserAction)
	collect, err := redis.Int64s(conn.Receive())
	userAction.Collect = collect
	comment, err := redis.Int64s(conn.Receive())
	userAction.Comment = comment
	like, err := redis.Int64s(conn.Receive())
	userAction.Like = like
	approve, err := redis.Int64s(conn.Receive())
	userAction.Approve = approve
	browse, err := redis.Int64s(conn.Receive())
	userAction.Browse = browse
	if err != nil {
		golog.Error(err)
	}
	return userAction
}

func getActionCount(refId uint64, kind int8) *ov.ActionCount {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	id := strconv.FormatUint(refId, 10)
	key := strings.Join([]string{IndexKind[kind], id, "Action", "Count"}, "_")
	conn.Send("SELECT", kind)
	conn.Send("HGETALL", key)
	conn.Flush()
	conn.Receive()
	actionCount := new(ov.ActionCount)
	action, err := redis.Int64Map(conn.Receive())
	actionCount.CollectCount = action[IndexAction[actionCollect]]
	actionCount.LikeCount = action[IndexAction[actionLike]]
	actionCount.ApproveCount = action[IndexAction[actionApprove]]
	actionCount.BrowseCount = action[IndexAction[actionBrowse]]
	actionCount.CommentCount = action[IndexAction[actionComment]]
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
	like.UserID = ctx.Values().Get("userID").(uint64)

	/*	var count int
		err := initialize.DB.Where(&like).First(&like).Count(&count).Error
		if count > 0 {
			if like.GETStatus()==0{
				err =initialize.DB.Model(&like).Update("status",1).Error
				setCountToRedis(like.UserID, like.RefID, KindIndex[like.Kind], actionLike, 1)
				common.Response(ctx, "成功", e.SUCCESS)
				return
			}else {
				err =initialize.DB.Model(&like).Update("status",0).Error
				setCountToRedis(like.UserID, like.RefID, KindIndex[like.Kind], actionLike, -1)
				common.Response(ctx, "成功取消", e.Sub)
				return
			}

		}*/

	if res, err := reUpdateStatus(&like, actionLike); err == nil {
		common.Response(ctx, "成功", res)
		return
	}

	like.Status = 1
	like.CreatedAt = time.Now()
	err := initialize.DB.Create(&like).Error
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "喜欢失败", e.ERROR)
		return
	}

	setCountToRedis(like.UserID, like.RefID, KindIndex[like.Kind], actionLike, 1)

	common.Response(ctx, "成功", e.SUCCESS)
}

func updateStatus(kind Action, action int8) (int, error) {
	var count int

	err := initialize.DB.Where(kind).First(kind).Count(&count).Error
	if count > 0 {
		if kind.GETStatus() == 0 {
			err = initialize.DB.Model(kind).Update("status", 1).Error
			setCountToRedis(kind.GETUserID(), kind.GETRefID(), kind.GETKind(), action, 1)
			return e.SUCCESS, err
		} else {
			err = initialize.DB.Model(kind).Update("status", 0).Error
			setCountToRedis(kind.GETUserID(), kind.GETRefID(), kind.GETKind(), action, -1)
			return e.Sub, err
		}

	}
	return 0, err
}

func reUpdateStatus(kind interface{}, action int8) (int, error) {
	var count int
	err := initialize.DB.Where(kind).First(kind).Count(&count).Error
	v := reflect.ValueOf(kind).Elem()
	status := v.FieldByName("Status").Interface().(uint8)
	userID := v.FieldByName("UserID").Uint()
	refID := v.FieldByName("RefID").Uint()
	Kind := v.FieldByName("Kind").String()
	if count > 0 {
		if status == 0 {
			err = initialize.DB.Model(kind).Update("status", 1).Error
			setCountToRedis(userID, refID, KindIndex[Kind], action, 1)
			return e.SUCCESS, err
		} else {
			err = initialize.DB.Model(kind).Update("status", 0).Error
			setCountToRedis(userID, refID, KindIndex[Kind], action, -1)
			return e.Sub, err
		}

	}
	return 0, err
}

func Approve(ctx iris.Context) {
	type Approve struct {
		RefID uint64 `json:"ref_id"`
		Kind  string `json:"kind"`
	}
	var approve Approve
	if err := ctx.ReadJSON(&approve); err != nil {
		common.Response(ctx, "参数无效")
		return
	}
	userID := ctx.Values().Get("userID").(uint64)
	var count int
	initialize.DB.Table(approve.Kind+"_approve_user").Where(approve.Kind+"_id = ? AND user_id = ?", approve.RefID, userID).Count(&count)
	if count > 0 {
		initialize.DB.Exec("DELETE FROM " + approve.Kind + "_approve_user WHERE " + approve.Kind + "_id =" +
			strconv.FormatUint(approve.RefID, 10) + " AND user_id = " + strconv.FormatUint(userID, 10))
		setCountToRedis(userID, approve.RefID, KindIndex[approve.Kind], actionApprove, -1)
		common.Response(ctx, "成功", e.Sub)
		return
	}

	err := initialize.DB.Exec("INSERT INTO " + approve.Kind + "_approve_user VALUES (" +
		strconv.FormatUint(approve.RefID, 10) + "," + strconv.FormatUint(userID, 10) + ")").Error
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "点赞失败", e.ERROR)
		return
	}

	setCountToRedis(userID, approve.RefID, KindIndex[approve.Kind], actionApprove, 1)

	common.Response(ctx, "成功", e.SUCCESS)
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

	userID := ctx.Values().Get("userID").(uint64)
	var count int
	initialize.DB.Model(&model.Favorites{}).Where("user_id =? AND id in (?)", userID, fc.FavoritesIDs).Count(&count)
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
		err = initialize.DB.Create(&Collection{RefID: fc.RefID, Kind: fc.Kind, UserID: userID, FavoritesID: v, Status: 1, CreatedAt: time.Now()}).Error
	}
	if err != nil {
		golog.Error(err)
		common.Response(ctx, "收藏失败", e.ERROR)
		return
	}

	setCountToRedis(userID, fc.RefID, KindIndex[fc.Kind], actionCollect, 1)

	common.Response(ctx, "收藏成功", e.SUCCESS)
}

func DelCollection(ctx iris.Context) {

}

func GetFavorite(ctx iris.Context) {
	id := ctx.Values().Get("userID").(uint64)
	var favorites []Favorites
	initialize.DB.Where("user_id=?", id).Find(&favorites)
	common.Response(ctx, favorites)
}

func AddFavorite(ctx iris.Context) {
	userID := ctx.Values().Get("userID").(uint64)

	var f Favorites
	if err := ctx.ReadJSON(&f); err != nil {
		common.Response(ctx, "参数无效")
		return
	}
	f.UserID = userID
	var count int
	initialize.DB.Model(&model.Favorites{}).Where(&f).Count(&count)
	if count > 0 {
		common.Response(ctx, "收藏夹已存在")
		return
	}
	f.CreatedAt = time.Now()
	initialize.DB.Create(&f)
	common.Response(ctx, f, "添加成功", e.SUCCESS)
}
