package main

import "fmt"

func main()  {
	s:=[]int{1,2,3,4,5,6}
	s =s[0:1:2]//这是什么操作
	fmt.Printf("%v",s)
}
