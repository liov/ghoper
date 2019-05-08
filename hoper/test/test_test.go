package test

import (
	"hoper/client/controller/cron"
	"testing"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/3
 * @description：
 */
func TestAll(t *testing.T) {
	cron.UserRedisToDB()
}
