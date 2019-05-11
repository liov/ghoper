package main

import (
	"fmt"
	"time"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/3
 * @description：
 */

func main() {
	var StartTime = time.Now().Unix()
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().Unix() - StartTime)
}
