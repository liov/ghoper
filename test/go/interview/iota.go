package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	starvationThresholdNs = 1e6
)

func main()  {
	fmt.Println(mutexWaiterShift)
}
