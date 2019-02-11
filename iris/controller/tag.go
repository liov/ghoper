package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
	"service/controller/common"
	"service/initialize"
	"service/model"
	"service/utils"
	"strconv"
	"time"
)

type Tag struct {
	Description string    `gorm:"type:varchar(100)" json:"description"`
	Name        string    `gorm:"type:varchar(10);primary_key" json:"name"`
	Count       uint      `gorm:"default:0" json:"count"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uint      `json:"created_by"`
}

func GetTags(c iris.Context) {
	args := c.QueryArgs()
	pageNo, _ := strconv.Atoi(utils.ToSting(args.Peek("pageNo")))
	pageSize, _ := strconv.Atoi(utils.ToSting(args.Peek("pageSize")))

	var tags []Tag

	err := initialize.DB.Select("name").
		Order("count desc").Limit(pageSize).Offset(pageNo).Find(&tags).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	common.Response(c, tags)
}

func GetTagTotal(maps interface{}) (count int) {
	initialize.DB.Model(&model.Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) *Tag {
	var tag Tag
	initialize.DB.Select("name,count").Where("name = ?", name).First(&tag)
	if tag.Name != "" {
		return &tag
	}

	return nil
}

func AddTag(c iris.Context) bool {
	args := c.QueryArgs()
	name := utils.ToSting(args.Peek("name"))
	user := c.UserValue("user").(User)
	initialize.DB.Create(&Tag{
		Name:      name,
		CreatedBy: user.ID,
	})

	return true
}

/*func ExistTagByID(id int) bool {
	var tag model.Tag
	initialize.DB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}*/

func DeleteTag(id int) bool {
	initialize.DB.Where("id = ?", id).Delete(&model.Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	initialize.DB.Model(&model.Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllTag() bool {
	initialize.DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&model.Tag{})

	return true
}
