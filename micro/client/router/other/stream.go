package other

import (
	"fmt"
	"github.com/kataras/iris"
	"io"
	"time"
)

func Stream(ctx iris.Context) {
	ctx.ContentType("text/html")
	ctx.Header("X-Accel-Buffering", "no") //nginx的锅必须加
	ctx.Header("Transfer-Encoding", "chunked")
	i := 0
	ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
	ctx.StreamWriter(func(w io.Writer) bool {
		fmt.Fprintf(w, "Message number %d<br>", ints[i])
		time.Sleep(500 * time.Millisecond) // simulate delay.
		if i == len(ints)-1 {
			return false //关闭并刷新
		}
		i++
		return true //继续写入数据
	})
}
