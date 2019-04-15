package noinit

import (
	"fmt"
	"hoper/model/ov"
	"reflect"
	"testing"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/11
 * @description：
 */

func TestType(t *testing.T) {
	var comments = func(num int) interface{} {
		switch num {
		case 1:
			return make([]ov.ArticleComment, 0, 5)

		case 2:
			return make([]ov.MomentComment, 0, 5)

		case 3:
			return make([]ov.DiaryComment, 0, 5)

		case 4:
			return make([]ov.DiaryBookComment, 0, 5)
		}
		return nil
	}(1)

	var c []ov.ArticleComment
	tt := reflect.TypeOf(comments)
	tc := reflect.TypeOf(c)
	fmt.Println(tt.Name(), tc.Name())
	fmt.Println(tt.Elem().Name(), tc.Elem().Name())

}
