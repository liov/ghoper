package initialize

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var engine *xorm.Engine

func initializeXorm() error {
	var err error
	url := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		DatabaseSettings.Host, DatabaseSettings.User, DatabaseSettings.Database, DatabaseSettings.Password)
	engine, err = xorm.NewEngine("postgres", url)
	if err != nil {
		return err
	}
	return nil
}
