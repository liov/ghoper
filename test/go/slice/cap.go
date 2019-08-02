package main

import "fmt"

/**
debug了一下，跳不进去runtime
最后总结了一下其实没有那么复杂
if oldCap*2 <= oldLen + addLen {
	newCap = oldCap*2
} else {
	newCap = oldLen + addLen
	if newCap & 1 == 1{	//判断单双
		newCap = newCap	+ 1 //单数 + 1
	}
}
//例:原切片len：7,cap:12,添加18个元素，新切片len:25,cap:26, 26 = 18 + 7 +1
//例:原切片len：7,cap:12,添加19个元素，新切片len:26,cap:26, 26 = 19 + 7
*/
func main() {
	var s1 []int
	fmt.Println(cap(s1)) //0
	s1 = append(s1, 0)   //空切片append容量变为1，然后开始二倍扩容
	fmt.Println(len(s1),cap(s1)) //1
	s1 = append(s1, 0)
	fmt.Println(len(s1),cap(s1))     //2
	s1 = append(s1, 0, 0, 0) //目前没搞清楚这里怎么扩容,像是添加元素和原cap比较，大的二倍
	fmt.Println(len(s1),cap(s1))     //6
	s1 = append(s1, 0, 0)
	fmt.Println(len(s1),cap(s1)) //12
	var s2 =[]int{0, 0,0,0,0,0,0, 0,0,0,0,0,0, 0,0,0,0,0}
	fmt.Println("s2:",len(s2),cap(s2))
	//扩容取决于s2的数量，如果s1的容量*2刚好可以放下新切片,那么扩容就是简单*2
	//如果不可以放下newCap = oldCap*2 + (h:=((addLen + oldLen) - oldCap*2))&0 == 1?h + 1 ：h,注：括号内赋值是java的语法
	//意思是在原容量扩容二倍的基础上,超出部分长度是单数则加(超出部分长度+1),双数则直接相加
	//即除初始化外，保证自扩容的切片容量是二的倍数
	//例:原切片len：7,cap:12,添加18个元素，新切片len:25,cap:26, 26 = 12*2 + (((18+7) - 12*2) + 1)
	//例:原切片len：7,cap:12,添加19个元素，新切片len:26,cap:26, 26 = 12*2 + ((19+7) - 12*2)
	s1 = append(s1,s2... )
	fmt.Println(len(s1),cap(s1)) //32
	var s3 = make([]int,3,5)
	fmt.Println(len(s3),cap(s3)) //5
	s3 = append(s3, 0, 0,0)
	fmt.Println(len(s3),cap(s3)) //10
}
