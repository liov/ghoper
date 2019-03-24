package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string
	if s == "" {
		print("初始化为空")
	}

	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(", ")
	buffer.WriteString("world")

	fmt.Print(buffer.String())
}
