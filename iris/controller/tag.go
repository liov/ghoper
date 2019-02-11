package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"service/controller/common"
	"service/initialize"
	"service/model"
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

	pageNo, _ := strconv.Atoi(c.URLParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.URLParam("pageSize"))

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

	name := c.URLParam("name")
	user := c.GetViewData()["user"].(User)
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
