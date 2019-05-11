package mixture

import (
	"test/main/struct/mixture/a"
	"test/main/struct/mixture/b"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/1
 * @description：
 */

type A struct {
	a.A
	b.A
}

type B struct {
	b.B
}

func main() {
	var b B
	b.I = 1
	b.A.I = 1
}
