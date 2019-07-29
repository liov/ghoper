package main

import (
	"fmt"
	"time"
)

func main() {
	for{
		select {
		case <-time.After(time.Second*2):
			fmt.Println("2秒定时器")
		}
	}
}
