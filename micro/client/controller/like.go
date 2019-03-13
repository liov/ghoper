package controller

import (
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/initialize"
	"time"
)

type Collection struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `gorm:"type:varchar(20)" json:"name"`
	User      *User      `json:"user"`
	UserID    uint       `json:"user_id"`
	Count     uint       `json:"count"`
	Articles  []Article  `gorm:"many2many:article_collection" json:"articles"`
	Moments   []Moment   `gorm:"many2many:moment_collection" json:"moments"`
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
	var collection []Collection
	initialize.DB.Where("user_id=?", id).Find(&collection)
	common.Response(ctx, collection)
}
