package po

import (
	"hoper/model/new/vo"
)

type Resume struct {
	vo.Resume
	vo.ResumeOwn
	vo.CUDTime
}

type Education struct {
	vo.Education
	vo.CUDTime
}

type Work struct {
	vo.Work
	vo.CUDTime
}
