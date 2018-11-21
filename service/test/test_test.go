package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

var where = map[string]interface{} {
"product_id = ?" : 10,
"user_id = ?" : 1232 ,
}


func TestTs(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "123456")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
	Ts()
	fmt.Println(where)


}

