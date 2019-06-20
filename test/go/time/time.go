package main

import (
	"log"
	"time"
)
//windows的Sleep，2ms算可以了？
func main()  {
	now:=time.Now().UnixNano()
	start:=now
	for{
		t:=0
		for start >=now{
			t++
			time.Sleep(1)
			now = time.Now().UnixNano()
		}
		start = now
		if t > 1 {
			log.Printf("t:%d",t)
		}
	}
}
