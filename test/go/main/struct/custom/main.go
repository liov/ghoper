package main

import (
	"bufio"
	"fmt"
	"io"
)

type A struct {
	I int
	S string
}

type B struct {
	J int
	S string
}

func (b *B) MB() {
	b.J = 1
}

type My struct {
	a A
	B
	int
	f  func(int) int
	SS string
}

func (m My) print() { // 值方法
	fmt.Println("Human:", m)
}

func (m My) MV() { // 值方法
	m.S = "s"
}

func (m *My) MX() { // 值方法
	m.S = "s"
}

type Context interface {
	io.Writer
	Print()
}

type Reader bufio.Reader

func (r *Reader) MC() {
	fmt.Println(r)
}

type Reader2 = Reader

func (r *Reader2) MR() {
	fmt.Println(r)
}

type F func(int)

func (f F) exe() {
	f(1)
}

func (f *F) Exe() {
	f.exe()
}

func (f *F) Ex() {
	ff := *f
	ff(1)
}

func main() {

	my := new(My)
	fmt.Println(my)
	my.MV()
	fmt.Println(my)
	my.MX()
	fmt.Println(my)
	my.MB()
	fmt.Println(my)
	r := new(Reader)
	fmt.Println(r)

	f := new(F)
	fmt.Println(f)
	var ff F = func(a int) {
		fmt.Println(a)
	}
	ff.exe()
	ff.Exe()
}

/*接收器不能是一个指针类型，但是它可以是任何其他允许类型的指针。
type MyInt int

type Q *MyInt

func (q Q) print() { // invalid receiver type Q (Q is a pointer type)
	fmt.Println("Q:", q)
}*/
