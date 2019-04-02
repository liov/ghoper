package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"hoper/client/controller/cachekey"
	"hoper/client/controller/common"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/crm"
	"hoper/model/e"
	"hoper/model/ov"
	"hoper/utils"
	"hoper/utils/gredis"
	"hoper/utils/hlog"
	"strconv"
	"strings"
	"time"
)

type Moments struct {
	TopMoments    []ov.Moment `json:"top_moments"`
	NormalMoments []ov.Moment `json:"normal_moments"`
}

//其实这里就是可插拔的，把redis操作单独放进一个函数
func GetMoments(c iris.Context) {
	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	topNum, _ := strconv.Atoi(c.URLParam("t"))
	//l := list.New()
	topKey := cachekey.TopMoments
	normalKey := cachekey.Moments

	/*	var moments []Moment

		if gredis.Exists(key) {
			data, err := gredis.Get(key)
			if err != nil {
				logging.Info(err)
			} else {
				json.Unmarshal(data, &moments)
				common.Response(c, moments)
				return
			}
		}*/

	var moments Moments

	if moments, count := getRedisMoments(topKey, normalKey, pageNo, topNum); moments != nil {
		common.Res(c, iris.Map{"data": *moments,
			"count": count,
			"msg":   e.GetMsg(e.SUCCESS),
			"code":  e.SUCCESS})
		return
	}
	//gorm 的ORM 弃用，决定手写sql
	/*	err := initialize.DB.Preload("Tags").Preload("Mood").Order(order).Offset(pageNo).Limit(pageSize).Find(&moments).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}*/
	if pageNo == 0 {
		initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("name,moment_id")
		}).Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count").
			Where("sequence > ?", 0).Order("id desc").Find(&moments.TopMoments)
	}

	err := initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count").
		Where("sequence = ?", 0).Order("id desc").Limit(pageSize - len(moments.TopMoments)).
		Offset(pageNo*pageSize - topNum).Find(&moments.NormalMoments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	var count int
	initialize.DB.Find(&moments.NormalMoments).Count(&count)
	//原来想存进链表，但不知道链表怎么序列化
	/*	for i := 0; i < len(moments); i++ {
			l.PushBack(moments[i])
		}

		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}*/

	//排序这种事是交给前端还是后端呢，给前端吧，代码多，给后端吧，怕效率不行
	/*
		//为了性能考虑，手写sql，联表查询的结果不组装对象，全部丢给前端，让前端去处理
		momentSql := "SELECT id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count,permission FROM moment WHERE status=0 ORDER BY desc_flag desc, created_at desc LIMIT ? OFFSET ?"
		initialize.DB.Raw(momentSql,pageSize,pageNo).Scan(&moments)
		tagsSql :="SELECT name FROM tag INNER JOIN moment_tag ON moment_tag.tag_name = tag.name WHERE (moment_tag.moment_id IN ('7','6','5','4','3','2','1')) AND status=0"
		type MomentTag struct {
			MomentID uint64 `json:"moment_id"`
			TagName string `json:"tag_name"`
		}
		var tags []MomentTag
		//循环遍历组装对象
		initialize.DB.Raw(momentSql,pageSize,pageNo).Scan(&tags)
		for mi, mv := range moments {
			for ti,tv := range tags{

			}
		}
	*/

	if len(moments.NormalMoments) == 0 && len(moments.NormalMoments) == 0 {
		common.Res(c, iris.Map{"data": moments,
			"count": count,
			"msg":   e.GetMsg(e.SUCCESS),
			"code":  e.SUCCESS})
		return
	}

	setRedisMoments(topKey, normalKey, moments, count)

	if moments, count := getRedisMoments(topKey, normalKey, pageNo, topNum); moments != nil {
		common.Res(c, iris.Map{"data": *moments,
			"count": count,
			"msg":   e.GetMsg(e.SUCCESS),
			"code":  e.SUCCESS})
		return
	}
}

func getRedisMoments(topKey string, normalKey string, pageNo int, topNum int) (*Moments, int) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	var moments Moments
	/*	if exist, err := redis.Bool(conn.Do("EXISTS", topKey)); !exist || err != nil {
		return nil
	}*/
	if exist, err := redis.Bool(conn.Do("EXISTS", normalKey)); !exist || err != nil {
		return nil, 0
	}

	if pageNo == 0 {
		topData, _ := redis.Strings(conn.Do("LRANGE", topKey, 0, -1))
		for mi, mv := range topData {
			if mv != "" {
				var moment ov.Moment
				utils.Json.UnmarshalFromString(mv, &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				moments.TopMoments = append(moments.TopMoments, moment)
				data, _ := utils.Json.MarshalToString(&moment)
				conn.Do("LSET", topKey, mi, data)
			} else {
				moments.TopMoments = append(moments.TopMoments, ov.Moment{})
			}
		}
		topNum = len(moments.TopMoments)
	}

	start := pageNo*model.PageSize - topNum
	if start < 0 {
		start = 0
	}

	if pageNo > 0 {
		topNum = 0
	}

	data, _ := redis.Strings(conn.Do("LRANGE", normalKey, start, start+model.PageSize-topNum-1))
	for mi, mv := range data {
		if mv != "" {
			var moment ov.Moment
			utils.Json.UnmarshalFromString(mv, &moment)
			moment.BrowseCount = moment.BrowseCount + 1
			moments.NormalMoments = append(moments.NormalMoments, moment)
			data, _ := utils.Json.MarshalToString(&moment)
			conn.Do("LSET", normalKey, mi+start, data)
		} else {
			moments.NormalMoments = append(moments.NormalMoments, ov.Moment{})
		}
	}

	if moments.NormalMoments == nil && moments.TopMoments == nil {
		return nil, 0
	}

	count, _ := redis.Int(conn.Do("GET", "Moment_List_Count"))
	return &moments, count
}

func setRedisMoments(topKey string, normalKey string, moments Moments, count int) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")

	if len(moments.TopMoments) > 0 {
		for _, mv := range moments.TopMoments {
			mv.BrowseCount = mv.BrowseCount + 1
			//mv.Index = mi
			value, _ := utils.Json.MarshalToString(mv)
			conn.Send("RPUSH", topKey, value)

		}
	}

	for _, mv := range moments.NormalMoments {
		mv.BrowseCount = mv.BrowseCount + 1
		//mv.Index = mi
		value, _ := utils.Json.MarshalToString(mv)
		conn.Send("RPUSH", normalKey, value)
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		return err
	}

	conn.Do("SET", "Moment_List_Count", strconv.Itoa(count))
	/*	_, err := conn.Do("EXPIRE", topKey, time)
		if err != nil {
			return err
		}*/
	return nil
}
func GetMoment(c iris.Context) {

	top := c.URLParam("t")
	index := c.URLParam("index")

	userID := c.Values().Get("userID").(uint64)

	if moment := getRedisMoment(top, index); moment != nil {
		if moment.UserID == userID {
			common.Response(c, *moment, "belong")
		} else {
			common.Response(c, *moment)
		}

		return
	}

	var moment ov.Moment

	id, err := c.Params().GetUint64("id")

	err = initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count,permission").
		Where("id = ?", id).First(&moment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	moment.BrowseCount = moment.BrowseCount + 1

	if moment.UserID == userID {
		common.Response(c, moment, "belong")
	} else {
		common.Response(c, moment)
	}

	saveErr := initialize.DB.Save(&moment).Error

	if saveErr != nil {
		hlog.Info("保存失败")
	}

}

func getRedisMoment(top string, index string) *ov.Moment {

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	//1代表是置顶，0代表不是
	if top != "0" {
		if gredis.Exists(cachekey.TopMoments) {
			data, err := conn.Do("LINDEX", cachekey.TopMoments, index)
			if err != nil {
				golog.Error(err)
			}
			if data != "" {
				var moment ov.Moment
				utils.Json.Unmarshal(data.([]byte), &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				data, err = utils.Json.MarshalToString(moment)
				_, err = conn.Do("LSET", cachekey.TopMoments, index, data)
				if err != nil {
					golog.Error(err)
				}
				return &moment
			} else {
				return nil
			}
		}
	} else {
		if gredis.Exists(cachekey.Moments) {
			data, err := conn.Do("LINDEX", cachekey.Moments, index)
			if err != nil {
				hlog.Info(err)
			}
			if data != "" {
				var moment ov.Moment
				utils.Json.Unmarshal(data.([]byte), &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				data, err = utils.Json.MarshalToString(moment)
				_, err = conn.Do("LSET", cachekey.Moments, index, data)
				if err != nil {
					hlog.Error(err)
				}
				return &moment
			} else {
				return nil
			}
		}
	}
	return nil
}

//isDel 0否 1是
func historyMoment(c iris.Context, isDel uint8) (*model.Moment, error) {

	//获取文章ID
	id := c.Params().GetUint64Default("id", 0)

	var moment model.Moment

	if err := initialize.DB.Preload("Tags").First(&moment, id).Error; err != nil {
		common.Response(c, "无效的版块id")
		return nil, err
	}

	momentHistory := crm.MomentHistory{
		//EverCreatedAt : moment.CreatedAt,
		CreatedAt:   time.Now(),
		MomentID:    moment.ID,
		ModifyTimes: moment.ModifyTimes + 1,
		DeleteFlag:  isDel,
		Content:     moment.Content,
		ImageUrl:    moment.ImageUrl,
		MoodName:    moment.MoodName,
		//Comments:    moment.Comments,
		UserID: moment.UserID,
	}

	saveErr := initialize.DB.Create(&momentHistory).Error

	for _, v := range moment.Tags {
		momentHistoryTag := crm.MomentHistoryTag{MomentHistoryID: momentHistory.ID, TagName: v.Name}
		initialize.DB.Create(&momentHistoryTag)
	}

	if saveErr != nil {
		hlog.Info("保存历史失败")
	}

	return &moment, nil
}

func EditMoment(c iris.Context) {

	moment, _ := historyMoment(c, 0)
	//userID := c.Values().Get("userID").(uint64)
	var newMoment model.Moment
	if err := c.ReadJSON(&newMoment); err != nil {
		common.Response(c, "参数无效")
		return
	}
	//moment.CreatedAt = momentHistory.EverCreatedAt
	//newMoment.ID = moment.ID

	//newMoment.Mood = model.Mood{Name: newMoment.MoodName}

	newMoment.ModifyTimes = moment.ModifyTimes + 1

	nowTime := time.Now()

	if newMoment.MoodName != "" {
		if mood := ExistMoodByName(newMoment.MoodName); mood != nil {
			setFlagCountToRedis(flagMood, newMoment.MoodName, 1)
		}
	}

	var tmpTags []ov.Tag
	if len(newMoment.Tags) > 0 {

		initialize.DB.Where("moment_id = ?", moment.ID).Delete(model.MomentTag{})
		var tagStringSlice []string
		for _, v := range moment.Tags {
			//tagStringSlice = append(tagStringSlice, v.Name)
			setFlagCountToRedis(flagTag, v.Name, -1)
		}

		tagString := strings.Join(tagStringSlice, ",")
		for _, v := range newMoment.Tags {
			tmpTags = append(tmpTags, ov.Tag{Name: v.Name, Description: v.Description})
			if !strings.Contains(tagString, v.Name) {
				if ExistTagByName(&v, moment.UserID) {
					setFlagCountToRedis(flagTag, v.Name, 1)
				}
			}
			momentTag := model.MomentTag{MomentID: moment.ID, TagName: v.Name}
			initialize.DB.Create(&momentTag)
		}
		newMoment.Tags = nil
	} else {
		for _, v := range moment.Tags {
			tmpTags = append(tmpTags, ov.Tag{Name: v.Name, Description: v.Description})
		}
	}
	newMoment.UpdatedAt = &nowTime

	//再留个坑
	moment.Tags = nil
	initialize.DB.Model(&moment).Updates(&newMoment)

	/*	saveErr := initialize.DB.Save(&moment).Error

		if saveErr != nil {
			logging.Info("修改失败")
			return
		}
	*/

	topNum := c.URLParam("t")
	index := c.URLParam("index")

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	redisMoment := ov.Moment{
		ID:        moment.ID,
		CreatedAt: moment.CreatedAt,
		Content:   moment.Content,
		ImageUrl:  moment.ImageUrl,
		Mood:      ov.Mood{Name: moment.MoodName},
		MoodName:  moment.MoodName,
		Tags:      tmpTags,
		UserID:    moment.UserID,
		ActionCount: ov.ActionCount{
			BrowseCount:  moment.BrowseCount,
			CommentCount: moment.CommentCount,
			CollectCount: moment.CollectCount,
			LikeCount:    moment.LikeCount,
		},
		Permission: moment.Permission,
	}
	//topNum
	if topNum != "0" {
		if gredis.Exists(cachekey.TopMoments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", cachekey.TopMoments, index, data)
			if err != nil {
				hlog.Error(err)
			}
		}
	} else {
		if gredis.Exists(cachekey.Moments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", cachekey.Moments, index, data)
			if err != nil {
				hlog.Error(err)
			}
		}
	}

	common.Response(c, "修改成功")
}

func DeleteMoment(c iris.Context) {

	historyMoment(c, 1)

	id := c.Params().GetUint64Default("id", 0)

	nowTime := time.Now()
	initialize.DB.Model(&model.Moment{ID: id}).Updates(&model.Moment{DeletedAt: &nowTime})

	topNum := c.URLParam("t")
	index := c.URLParam("index")

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if topNum != "0" {
		if gredis.Exists(cachekey.TopMoments) {
			_, err := conn.Do("LSET", cachekey.TopMoments, index, "")
			if err != nil {
				hlog.Error(err)
			}
		}
	} else {
		if gredis.Exists(cachekey.Moments) {
			_, err := conn.Do("LSET", cachekey.Moments, index, "")
			if err != nil {
				hlog.Error(err)
			}
		}
	}
	common.Response(c, "删除成功")

}
