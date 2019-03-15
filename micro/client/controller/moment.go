package controller

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"hoper/client/controller/common"
	"hoper/client/controller/common/e"
	"hoper/client/controller/common/gredis"
	"hoper/client/controller/common/logging"
	"hoper/initialize"
	"hoper/model"
	"hoper/utils"

	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//DTO
type Moment struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Content      string    `gorm:"type:varchar(500)" json:"content"`
	ImageUrl     string    `gorm:"type:varchar(100)" json:"image_url"` //图片
	Mood         Mood      `gorm:"foreignkey:MoodName;association_foreignkey:Name" json:"mood"`
	MoodName     string    `gorm:"type:varchar(20)" json:"mood_name"`
	Tags         []Tag     `gorm:"many2many:moment_tag;foreignkey:ID;association_foreignkey:Name" json:"tags"`
	User         User      `json:"user"`
	UserID       uint      `json:"user_id"`
	BrowseCount  uint      `json:"browse_count"`                                       //浏览
	CommentCount uint      `json:"comment_count"`                                      //评论
	CollectCount uint      `json:"collect_count"`                                      //收藏
	ApproveCount uint      `gorm:"default:0" json:"approve_count"`                     //点赞
	LikeCount    uint      `json:"like_count"`                                         //点赞
	Permission   uint8     `gorm:"type:smallint unsigned;default:0" json:"permission"` //查看权限
	//Index        int       `json:"index"`                                                //redis列表中排序
}

type Moments struct {
	TopMoments    []Moment `json:"top_moments"`
	NormalMoments []Moment `json:"normal_moments"`
}

//其实这里就是可插拔的，把redis操作单独放进一个函数
func GetMoments(c iris.Context) {
	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	topNum, _ := strconv.Atoi(c.URLParam("t"))
	//l := list.New()
	topKey := gredis.TopMoments
	normalKey := gredis.Moments

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
			MomentID uint `json:"moment_id"`
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
				var moment Moment
				utils.Json.UnmarshalFromString(mv, &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				moments.TopMoments = append(moments.TopMoments, moment)
				data, _ := utils.Json.MarshalToString(&moment)
				conn.Do("LSET", topKey, mi, data)
			} else {
				moments.TopMoments = append(moments.TopMoments, Moment{})
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
			var moment Moment
			utils.Json.UnmarshalFromString(mv, &moment)
			moment.BrowseCount = moment.BrowseCount + 1
			moments.NormalMoments = append(moments.NormalMoments, moment)
			data, _ := utils.Json.MarshalToString(&moment)
			conn.Do("LSET", normalKey, mi+start, data)
		} else {
			moments.NormalMoments = append(moments.NormalMoments, Moment{})
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

	user := c.Values().Get("user").(User)

	if moment := getRedisMoment(top, index); moment != nil {
		if moment.UserID == user.ID {
			common.Response(c, *moment, "belong")
		} else {
			common.Response(c, *moment)
		}

		return
	}

	/*
		if gredis.Exists(key) {
			data, err := gredis.Get(key)
			if err != nil {
				logging.Info(err)
			} else {
				json.Unmarshal(data, &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				gredis.Set(key, moment, 60)
				common.Response(c moment)
				return
			}
		}*/
	var moment Moment

	id, _ := strconv.Atoi(c.URLParam("id"))

	err := initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count,permission").
		Where("id = ?", id).First(&moment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	moment.BrowseCount = moment.BrowseCount + 1

	if moment.UserID == user.ID {
		common.Response(c, moment, "belong")
	} else {
		common.Response(c, moment)
	}

	saveErr := initialize.DB.Save(&moment).Error

	if saveErr != nil {
		logging.Info("保存失败")
		return
	}

}

func getRedisMoment(top string, index string) *Moment {

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	//1代表是置顶，0代表不是
	if top != "0" {
		if gredis.Exists(gredis.TopMoments) {
			data, err := conn.Do("LINDEX", gredis.TopMoments, index)
			if err != nil {
				logging.Error(err)
			}
			if data != "" {
				var moment Moment
				utils.Json.Unmarshal(data.([]byte), &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				data, err = utils.Json.MarshalToString(moment)
				_, err = conn.Do("LSET", gredis.TopMoments, index, data)
				if err != nil {
					logging.Error(err)
				}
				return &moment
			} else {
				return &Moment{}
			}
		}
	} else {
		if gredis.Exists(gredis.Moments) {
			data, err := conn.Do("LINDEX", gredis.Moments, index)
			if err != nil {
				logging.Info(err)
			}
			if data != "" {
				var moment Moment
				utils.Json.Unmarshal(data.([]byte), &moment)
				moment.BrowseCount = moment.BrowseCount + 1
				data, err = utils.Json.MarshalToString(moment)
				_, err = conn.Do("LSET", gredis.Moments, index, data)
				if err != nil {
					logging.Error(err)
				}
				return &moment
			} else {
				return &Moment{}
			}
		}
	}
	return nil
}

func AddMoment(c iris.Context) {

	user := c.Values().Get("user").(User)

	//Limit这个函数的封装呢，费了点功夫，之前的返回值想到用err，不过在sendErr这出了点问题，决定返回值改用string，这样是不规范的
	if limitErr := common.Limit(model.MomentMinuteLimit,
		model.MomentMinuteLimitCount,
		model.MomentDayLimit,
		model.MomentMinuteLimitCount, user.ID); limitErr != nil {
		common.Response(c, limitErr.Error(), e.TimeTooMuch)
		return
	}

	var moment Moment

	if err := c.ReadJSON(&moment); err != nil {
		logrus.WithFields(logrus.Fields{
			"model": "moment",
		}).Info(err.Error())
		common.Response(c, "参数无效")
		return
	}

	/*	moodName := moment.MoodName

		var mood model.Mood

		moodErr :=initialize.DB.Where("name = ?", moodName).Find(&mood).Error

		if moodErr != nil{
			mood.Name = moodName
			initialize.DB.Create(&mood)
			moment.Mood = mood
		} else {
			moment.Mood = mood
		}*/

	nowTime := time.Now()
	moment.CreatedAt = nowTime
	//moment.Mood = Mood{Name: moment.MoodName}

	if err := validationMoment(c, &moment); err != nil {
		return
	}
	moment.UserID = user.ID
	moment.BrowseCount = 1
	/*	moment.Status = model.ArticleVerifying
		moment.ModifyTimes = 0
		moment.ParentID = 0 */
	/*	user.Score = user.Score + model.ArticleScore
		user.ArticleCount = user.ArticleCount + 1

		if model.UserToRedis(user) != nil {
			common.SendErr(c,"error")
			return
		}*/
	moment.Content = strings.TrimSpace(moment.Content)
	var mood *Mood
	if mood = ExistMoodByName(moment.MoodName); mood != nil {
		initialize.DB.Model(mood).Update("count", mood.Count+1)
	} else {
		newMood := Mood{CreatedAt: nowTime, Name: moment.MoodName, Count: 1}
		initialize.DB.Create(&newMood)
		mood = &newMood
	}

	saveErr := initialize.DB.Create(&moment).Error

	for _, v := range moment.Tags {
		if tag := ExistTagByName(v.Name); tag != nil {
			initialize.DB.Model(tag).Update("count", tag.Count+1)
		} else {
			newTag := Tag{CreatedAt: nowTime, Name: v.Name, Count: 1}
			initialize.DB.Create(&newTag)
		}
		momentTag := model.MomentTag{MomentID: moment.ID, TagName: v.Name}
		initialize.DB.Create(&momentTag)
	}

	if saveErr != nil {
		common.Response(c, "创建出错")
		return
	}

	//var moments []model.Moment
	moment.User = user
	moment.Mood = *mood
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
	_, err := conn.Do("LPUSH", gredis.Moments, value)
	if err != nil {
		return
	}

	common.Response(c, "新建成功", e.SUCCESS)
}

func validationMoment(c iris.Context, moment *Moment) (err error) {

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

//isDel 0否 1是
func historyMoment(c iris.Context, isDel uint8) (*model.Moment, error) {

	//获取文章ID
	id := c.Params().GetUint64Default("id", 0)

	var moment model.Moment

	if err := initialize.DB.Preload("Tags").First(&moment, id).Error; err != nil {
		common.Response(c, "无效的版块id")
		return nil, err
	}

	nowTime := time.Now()

	momentHistory := model.MomentHistory{
		//EverCreatedAt : moment.CreatedAt,
		CreatedAt:   nowTime,
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
		momentHistoryTag := model.MomentHistoryTag{MomentHistoryID: momentHistory.ID, TagName: v.Name}
		initialize.DB.Create(&momentHistoryTag)
	}

	if saveErr != nil {
		logging.Info("保存历史失败")
	}

	return &moment, nil
}

func EditMoment(c iris.Context) {

	moment, _ := historyMoment(c, 0)

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
			initialize.DB.Model(mood).Update("count", mood.Count+1)
			oldMood := ExistMoodByName(moment.MoodName)
			initialize.DB.Model(oldMood).Update("count", oldMood.Count-1)
		} else {
			newMood := model.Mood{CreatedAt: nowTime, Name: newMoment.MoodName, Count: 1}
			initialize.DB.Create(&newMood)
		}
	}

	var tmpTags []Tag
	if len(newMoment.Tags) > 0 {

		initialize.DB.Where("moment_id = ?", moment.ID).Delete(model.MomentTag{})
		var tagStringSlice []string
		for _, v := range moment.Tags {
			//tagStringSlice = append(tagStringSlice, v.Name)
			tag := ExistTagByName(v.Name)
			initialize.DB.Model(tag).Update("count", tag.Count-1)
		}

		tagString := strings.Join(tagStringSlice, ",")
		for _, v := range newMoment.Tags {
			tmpTags = append(tmpTags, Tag{Name: v.Name, Description: v.Description})
			if !strings.Contains(tagString, v.Name) {
				if tag := ExistTagByName(v.Name); tag != nil {
					initialize.DB.Model(tag).Update("count", tag.Count+1)
				} else {
					newTag := model.Tag{CreatedAt: nowTime, Name: v.Name, Count: 1}
					initialize.DB.Create(&newTag)
				}
			}
			momentTag := model.MomentTag{MomentID: moment.ID, TagName: v.Name}
			initialize.DB.Create(&momentTag)
		}
		newMoment.Tags = nil
	} else {
		for _, v := range moment.Tags {
			tmpTags = append(tmpTags, Tag{Name: v.Name, Description: v.Description})
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

	redisMoment := Moment{
		ID:           moment.ID,
		CreatedAt:    moment.CreatedAt,
		Content:      moment.Content,
		ImageUrl:     moment.ImageUrl,
		Mood:         Mood{Name: moment.MoodName},
		MoodName:     moment.MoodName,
		Tags:         tmpTags,
		User:         User{},
		UserID:       moment.UserID,
		BrowseCount:  moment.BrowseCount,
		CommentCount: moment.CommentCount,
		CollectCount: moment.CollectCount,
		LikeCount:    moment.LikeCount,
		Permission:   moment.Permission,
	}
	//topNum
	if topNum != "0" {
		if gredis.Exists(gredis.TopMoments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", gredis.TopMoments, index, data)
			if err != nil {
				logging.Error(err)
			}
		}
	} else {
		if gredis.Exists(gredis.Moments) {
			data, err := utils.Json.MarshalToString(redisMoment)
			_, err = conn.Do("LSET", gredis.Moments, index, data)
			if err != nil {
				logging.Error(err)
			}
		}
	}

	common.Response(c, "修改成功")
}

func DeleteMoment(c iris.Context) {

	historyMoment(c, 1)

	id := c.Params().GetUint64Default("id", 0)

	nowTime := time.Now()
	initialize.DB.Model(&model.Moment{ID: uint(id)}).Updates(&model.Moment{DeletedAt: &nowTime})

	topNum := c.URLParam("t")
	index := c.URLParam("index")

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if topNum != "0" {
		if gredis.Exists(gredis.TopMoments) {
			_, err := conn.Do("LSET", gredis.TopMoments, index, "")
			if err != nil {
				logging.Error(err)
			}
		}
	} else {
		if gredis.Exists(gredis.Moments) {
			_, err := conn.Do("LSET", gredis.Moments, index, "")
			if err != nil {
				logging.Error(err)
			}
		}
	}
	common.Response(c, "删除成功")

}

func GetMomentsV2(c iris.Context) {
	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))
	//l := list.New()

	key := gredis.Moments + "_V2"

	var moments []Moment

	if moments, count, topCount := getRedisMomentsV2(key, pageNo, pageSize); moments != nil {
		common.Res(c, iris.Map{"data": moments,
			"count":     count,
			"top_count": topCount,
			"msg":       e.GetMsg(e.SUCCESS),
			"code":      e.SUCCESS})
		return
	}

	var count, topCount int
	err := initialize.DB.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("name,moment_id")
	}).Preload("User").Select("id,created_at,content,image_url,mood_name,user_id,browse_count,comment_count,collect_count,like_count").
		Order("sequence desc,id desc").Limit(pageSize).
		Offset(pageNo * pageSize).Find(&moments).Count(&count).Error
	err = initialize.DB.Model(Moment{}).Where("sequence = ?", 9).Count(&topCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	common.Res(c, iris.Map{"data": moments,
		"count":     count,
		"top_count": topCount,
		"msg":       e.GetMsg(e.SUCCESS),
		"code":      e.SUCCESS})

	setRedisMomentsV2(key, moments, count, topCount)

}

func getRedisMomentsV2(key string, pageNo int, PageSize int) ([]Moment, int, int) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	var moments []Moment
	if exist, err := redis.Bool(conn.Do("EXISTS", key)); !exist || err != nil {
		return nil, 0, 0
	}
	start := pageNo * PageSize

	data, _ := redis.Strings(conn.Do("LRANGE", key, start, start+PageSize-1))
	for mi, mv := range data {
		var moment Moment
		utils.Json.UnmarshalFromString(mv, &moment)
		moment.BrowseCount = moment.BrowseCount + 1
		moments = append(moments, moment)
		data, _ := utils.Json.MarshalToString(&moment)
		conn.Do("LSET", key, mi, data)

	}
	count, _ := redis.Int(conn.Do("GET", "Moment_List_Count"))
	topCount, _ := redis.Int(conn.Do("GET", "Moment_List_Top_Count"))
	return moments, count, topCount
}

func setRedisMomentsV2(key string, moments []Moment, count int, topCount int) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	for _, mv := range moments {
		mv.BrowseCount = mv.BrowseCount + 1
		//mv.Index = mi
		value, _ := utils.Json.MarshalToString(mv)
		_, err := conn.Do("RPUSH", key, value)
		if err != nil {
			return err
		}
	}
	conn.Do("SET", "Moment_List_Count", strconv.Itoa(count))
	conn.Do("SET", "Moment_List_Top_Count", strconv.Itoa(topCount))
	return nil
}
