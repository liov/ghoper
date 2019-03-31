package po

import "hoper/model/new/vo"

type Tag struct {
	vo.Tag
	vo.ExtraInfo
	vo.KindOwnCount
}

type Mood struct {
	vo.Mood
	vo.ExtraInfo
	vo.KindOwnCount
}

type Category struct {
	vo.Category
	vo.ExtraInfo
	vo.KindOwnCount
}
