package main

import "fmt"

func main() {
	message := "消息1"

	defer func() {
		fmt.Println("第一个defer：", message)
	}()

	message = "消息改变了"

	defer func(m string) {
		fmt.Println("第二个defer:", m)
	}(message)

	message = "消息2"
}
