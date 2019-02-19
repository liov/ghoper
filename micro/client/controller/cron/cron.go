package cron

import (
	"github.com/robfig/cron"
	"hoper/client/controller/crawler"
)

var funcMap = map[string]func(){}

var jobMap = map[string]cron.Job{}

func init() {
	funcMap["0 10 10 * * *"] = RedisToDB
	jobMap["0 0 2 * * *"] = RedisTo{}
	funcMap["0 0 0 * * *"] = crawler.MM131
}

// New 构造cron
func New() *cron.Cron {
	c := cron.New()
	for spec, cmd := range funcMap {
		c.AddFunc(spec, cmd)
	}
	for spec, cmd := range jobMap {
		c.AddJob(spec, cmd)
	}
	return c
}
