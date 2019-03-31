package po

import (
	"hoper/model/new/vo"
)

type User struct {
	vo.User
	Password string `gorm:"type:varchar(100)" json:"-"`
	vo.UserData
	vo.UserMore
	vo.UserNoPub
	vo.UserOwn
	vo.UserOwnCount
	vo.CUDTime
}

type Follow struct {
	vo.Follow
	vo.FollowNoPub
	vo.FollowStatus
}
