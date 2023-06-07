package controller

import (
	"hoper/utils/ulog"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"hoper/controller/common"
	"hoper/controller/credis"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/crm"
	"hoper/model/e"
	"hoper/model/ov"
	"hoper/utils"
	"hoper/utils/uredis"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/3/29
 * @description：
 */
//DTO

func AddMoment(c iris.Context) {

	user := c.Values().Get("user").(*User)

	//Limit这个函数的封装呢，费了点功夫，之前的返回值想到用err，不过在sendErr这出了点问题，决定返回值改用string，这样是不规范的
	if limitErr := common.Limit(model.MomentMinuteLimit,
		model.MomentMinuteLimitCount,
		model.MomentDayLimit,
		model.MomentMinuteLimitCount, user.ID); limitErr != nil {
		common.Response(c, limitErr.Error(), e.TimeTooMuch)
		return
	}

	var moment model.Moment

	if err := c.ReadJSON(&moment); err != nil {
		ulog.Error(err)
		common.Response(c, "参数无效")
		return
	}

	if utf8.RuneCountInString(moment.Content) > 500 {
		common.Response(c, "文章内容不能小于20个字")
		return
	}

	nowTime := time.Now()
	moment.CreatedAt = nowTime
	//moment.Mood = Mood{Name: moment.MoodName}

	if err := validationMoment(c, &moment); err != nil {
		return
	}
	moment.UserID = user.ID
	moment.BrowseCount = 1
	moment.Status = model.ArticleVerifying
	moment.ModifyTimes = 0
	moment.ParentID = 0
	user.Score = user.Score + model.ArticleScore
	user.ArticleCount = user.ArticleCount + 1

	if err := EditUserRedis(user); err != nil {
		ulog.Error(err)
	}
	moment.Content = strings.TrimSpace(moment.Content)

	if mood := ExistMoodByName(moment.MoodName); mood != nil {
		moment.Mood = *mood
		setFlagCountToRedis(flagTag, moment.MoodName, 1)
	}

	saveErr := initialize.DB.Create(&moment).Error

	for _, v := range moment.Tags {
		if ExistTagByName(&v, user.ID) {
			setFlagCountToRedis(flagTag, v.Name, 1)
		}
		momentTag := model.MomentTag{MomentID: moment.ID, TagName: v.Name}
		initialize.DB.Create(&momentTag)
	}

	if saveErr != nil {
		common.Response(c, "创建出错")
		return
	}

	//var moments []model.Moment
	moment.User = user.User

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	/*	if moment.DescFlag == 0 {
			value, _ := json.Marshal(moment)
			_, err := conn.Do("LPUSH", gredis.Moments, value)
			if err != nil {
				return
			}
		} else {
			value, _ := json.Marshal(moment)
			_, err := conn.Do("LPUSH", gredis.TopMoments, value)
			if err != nil {
				return
			}
		}*/

	value, _ := utils.Json.MarshalToString(moment)
	conn.Send("SELECT", kindMoment)
	conn.Send("LPUSH", credis.Moments, value)
	_, err := conn.Do("INCR", credis.Moments+"_Count")
	if err != nil {
		return
	}

	common.Response(c, "新建成功", e.SUCCESS)
}

func GetMoments(c iris.Context) {
	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	//l := list.New()
	userID := c.Values().Get("userID").(uint64)
	key := credis.Moments

	var moments []ov.Moment
	var userAction *UserAction
	if userID > 0 {
		userAction = GetRedisAction(strconv.FormatUint(userID, 10), kindMoment)
	}
	var count, topCount int64
	if moments, count, topCount = getRedisMoments(key, pageNo, pageSize); moments != nil {

		common.Res(c, iris.Map{"data": moments,
			"count":       count,
			"top_count":   topCount,
			"user_action": userAction,
			"msg":         e.GetMsg(e.SUCCESS),
			"code":        e.SUCCESS})
		return
	}

	err := initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Preload("User").Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count").
		Order("sequence desc,id desc").Limit(pageSize).
		Offset(pageNo * pageSize).Find(&moments).Count(&count).Error
	err = initialize.DB.Model(ov.Moment{}).Where("sequence = ?", 9).Count(&topCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if userID > 0 {
		GetRedisAction(strconv.FormatUint(uint64(userID), 10), kindMoment)
	}

	common.Res(c, iris.Map{"data": moments,
		"count":       count,
		"top_count":   topCount,
		"user_action": userAction,
		"msg":         e.GetMsg(e.SUCCESS),
		"code":        e.SUCCESS})

	setRedisMoments(key, moments, count, topCount)

}

func validationMoment(c iris.Context, moment *model.Moment) (err error) {

	err = &e.ValidtionError{Msg: "参数无效"}

	if moment.Content == "" || utf8.RuneCountInString(moment.Content) <= 0 {
		common.Response(c, "文章内容不能为空")
		return
	}

	if utf8.RuneCountInString(moment.Content) > model.MaxContentLen {
		msg := "文章内容不能超过" + strconv.Itoa(model.MaxContentLen) + "个字符"
		common.Response(c, msg)
		return
	}

	/*	if moment.Tags == nil || len(moment.Tags) <= 0 {
		SendErrJSON(c,"请选择标签")
		return
	}*/

	return nil
}

func getRedisMoments(key string, pageNo int, PageSize int) ([]ov.Moment, int64, int64) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	var moments []ov.Moment
	conn.Send("SELECT", kindMoment)
	if exist, err := redis.Bool(conn.Do("EXISTS", key)); !exist || err != nil {
		return nil, 0, 0
	}
	start := pageNo * PageSize

	data, _ := redis.Strings(conn.Do("LRANGE", key, start, start+PageSize-1))
	for _, mv := range data {
		var moment ov.Moment
		utils.Json.UnmarshalFromString(mv, &moment)
		conn.Send("HINCRBY", strings.Join([]string{IndexKind[kindMoment], strconv.FormatUint(moment.ID, 10), "Action", "Count"}, "_"), IndexAction[actionBrowse], 1)
		actionCount := getActionCount(moment.ID, kindMoment)
		actionCount.BrowseCount = actionCount.BrowseCount + 1
		moment.ActionCount = *actionCount
		moments = append(moments, moment)
	}
	conn.Do("")
	conn.Send("GET", "Moment_List_Count")
	conn.Send("GET", "Moment_List_Top_Count")
	conn.Flush()
	count, _ := redis.Int64(conn.Receive())
	topCount, _ := redis.Int64(conn.Receive())
	return moments, count, topCount
}

func setRedisMoments(key string, moments []ov.Moment, count int64, topCount int64) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("SELECT", kindMoment)
	for _, mv := range moments {
		mv.BrowseCount = mv.BrowseCount + 1
		//mv.Index = mi
		value, _ := utils.Json.MarshalToString(mv)
		err := conn.Send("RPUSH", key, value)
		if err != nil {
			return err
		}
	}
	conn.Send("SET", "Moment_List_Count", strconv.FormatInt(count, 10))
	conn.Do("SET", "Moment_List_Top_Count", strconv.FormatInt(topCount, 10))
	return nil
}

func GetMoment(c iris.Context) {

	index := c.URLParam("index")
	userID := c.Values().Get("userID").(uint64)
	var userAction *UserAction
	if userID > 0 {
		userAction = GetRedisAction(strconv.FormatUint(userID, 10), kindMoment)
	}
	if moment := getRedisMoment(index); moment != nil {

		common.Res(c, iris.Map{"data": moment,
			"user_action": userAction,
			"msg":         e.GetMsg(e.SUCCESS),
			"code":        e.SUCCESS})
		return
	}

	id, err := c.Params().GetUint64("id")
	var moment ov.Moment

	err = initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Preload("User").Where("id = ?", id).First(&moment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		ulog.Error(err)
		return
	}
	actionCount := getActionCount(moment.ID, kindMoment)
	moment.ActionCount = *actionCount
	common.Res(c, iris.Map{"data": moment,
		"user_action": userAction,
		"code":        e.SUCCESS})
}

func getRedisMoment(index string) *ov.Moment {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := credis.Moments
	conn.Send("SELECT", kindMoment)

	data, err := redis.String(conn.Do("LINDEX", key, index))
	var moment ov.Moment
	err = utils.Json.UnmarshalFromString(data, &moment)
	conn.Do("HINCRBY", strings.Join([]string{IndexKind[kindMoment], strconv.FormatUint(moment.ID, 10), "Action", "Count"}, "_"), IndexAction[actionBrowse], 1)
	actionCount := getActionCount(moment.ID, kindMoment)
	moment.ActionCount = *actionCount
	if err != nil {
		ulog.Error(err)
		return nil
	}
	return &moment

}

// isDel 0否 1是
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
		ulog.Info("保存历史失败")
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
		if uredis.Exists(credis.TopMoments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", credis.TopMoments, index, data)
			if err != nil {
				ulog.Error(err)
			}
		}
	} else {
		if uredis.Exists(credis.Moments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", credis.Moments, index, data)
			if err != nil {
				ulog.Error(err)
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
		if uredis.Exists(credis.TopMoments) {
			_, err := conn.Do("LSET", credis.TopMoments, index, "")
			if err != nil {
				ulog.Error(err)
			}
		}
	} else {
		if uredis.Exists(credis.Moments) {
			_, err := conn.Do("LSET", credis.Moments, index, "")
			if err != nil {
				ulog.Error(err)
			}
		}
	}
	common.Response(c, "删除成功")

}

func redisMoments(key string, model interface{}) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if exist, err := redis.Bool(conn.Do("EXISTS", key)); exist && err == nil {
		data, err := redis.Bytes(conn.Do("GET", key))
		if err != nil {
			ulog.Info(err)
			return err
		} else {
			utils.Json.Unmarshal(data, model)
			/*	for _, mv := range *moments {
						//瞬间是不需要设置缓存的，前端存储
						mkey := strings.Join([]string{
								e.CacheMoment,
								strconv.FormatUint(uint64(mv.ID),10),
							}, "_")

							mv.BrowseCount = mv.BrowseCount + 1

							_, err =conn.Do("SET", mkey, mv)
							_, err =conn.Do("EXPIRE", mkey, 60)

				}
				if err != nil {
					return err
				}
			*/
		}
	}
	return nil
}

func MomentRedisToDB() {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	conn.Send("SELECT", kindMoment)
	data, _ := redis.Strings(conn.Do("LRANGE", credis.Moments, 0, -1))
	for _, mv := range data {
		if mv != "" {
			var moment ov.Moment
			utils.Json.UnmarshalFromString(mv, &moment)
			initialize.DB.Model(&moment).UpdateColumns(ov.Moment{
				ActionCount: ov.ActionCount{
					CollectCount: moment.CollectCount,
					BrowseCount:  moment.BrowseCount, CommentCount: moment.CommentCount,
					LikeCount: moment.LikeCount},
			})
		}
	}
}
