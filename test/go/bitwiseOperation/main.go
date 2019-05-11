package main

import "fmt"

var a = 15
var b = 16

func main() {
	fmt.Println(a&1, b&1)
	fmt.Println(a|b, a|1)
	fmt.Println(a^b, a^a, a^b^b)
	fmt.Println(^a, ^b, 0^a)
	fmt.Println(a&^b, (a^b)&b) // 标志位操作 &^,清除标记位
	fmt.Println(a<<2, b<<2)
	fmt.Println(a>>2, b>>2)
}
