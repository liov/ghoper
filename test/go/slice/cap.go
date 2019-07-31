package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(cap(s1)) //0
	s1 = append(s1, 0)   //空切片append容量变为1，然后开始二倍扩容
	fmt.Println(cap(s1)) //1
	s1 = append(s1, 0)
	fmt.Println(cap(s1))     //2
	s1 = append(s1, 0, 0, 0) //目前没搞清楚这里怎么扩容,像是添加元素和原cap比较，大的二倍
	fmt.Println(cap(s1))     //6
	s1 = append(s1, 0, 0)
	fmt.Println(cap(s1)) //12
}
