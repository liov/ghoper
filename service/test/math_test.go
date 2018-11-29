package test

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestUpload(t *testing.T) {
	t1 := time.Now()
	for i := 0; i <= 10000; i++ {
		aaa(float64(i))
	}
	t2 := time.Now()
	fmt.Println("go time" + t2.Sub(t1).String())
}

func aaa(i float64) {
	a := i + 1
	b := 2.3
	s := "abcdefkkbghisdfdfdsfds"
	if a > b {
		a++
	} else {
		b += 1
	}

	if a == b {
		b += 1
	}

	c := a*b + a/b - a*a

	var d = s[0:strings.Index(s, "kkb")] + strconv.FormatFloat(c, 'E', -1, 64)
	fmt.Println(d)

}
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func stringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

func Test_ByteString(t *testing.T) {
	var x = []byte("Hello World!")
	var y = *(*string)(unsafe.Pointer(&x))
	var z = string(x)

	if y != z {
		t.Fail()
	}
}

func Benchmark_Normal(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func Benchmark_ByteString(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&x))
	}
}

type MyStruct struct {
	A int
	B string
	C *S
}

type S struct {
	E string
	D int
}

var sizeOfStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfStruct
	x.Cap = sizeOfStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}

func BenchmarkTs(b *testing.B) {
	a, _ := json.Marshal(&MyStruct{A: 10, B: "110"})
	var c MyStruct
	json.Unmarshal(a, &c)
	fmt.Println(a)
	for i := 0; i < b.N; i++ {
		MyStructToBytes(&MyStruct{})
	}
}

func BenchmarkTs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBytes("你好")
	}
}

func MyStructToBytes1(s interface{}) []byte {
	sizeOfStruct := reflect.TypeOf(s).Elem().Size()
	var x reflect.SliceHeader
	x.Len = (int)(sizeOfStruct)
	x.Cap = (int)(sizeOfStruct)
	x.Data = uintptr((*emptIntrtface)(unsafe.Pointer(&s)).word)
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct1(b []byte) unsafe.Pointer {

	return unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	)
}

type emptIntrtface struct {
	typ  *struct{}
	word unsafe.Pointer
}

func ToSting(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func sss(b []byte) string {
	return string(b)
}

func bbb(s string) []byte {
	return []byte(s)
}
