package rpc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func TestGo(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var a string
		var done bool
		go func() {

			a = fmt.Sprintf("%d_%s", i, "hello")

			done = true
		}()

		for !done {

		}

		fmt.Println("第", i, "次,a:", a)
	}
}

func TestRandom(t *testing.T) {
	for i := range random(100) {
		fmt.Println(i)
	}
}

func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func TestDuoXingShengChengQi(t *testing.T) {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		actState := initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func s() chan string {
	return make(chan string)
}

func ss() <-chan string {
	return make(chan string)
}

func TestChan(t *testing.T) {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func TestBiBao(t *testing.T) {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func TestPrint(t *testing.T) {
	x := decimalToAny(703, 26)
	flag := false
	var y []byte
	for i := range x {
		if !flag {
			if x[i] == '0' {
				y = append(y, 'Z')
				flag = true
			} else {
				if flag {
					if i == len(x)-1 && byte(x[i]) == '1' {
						break
					}
					if byte(x[i]) == '1' {
						y = append(y, 'Z')
					} else {
						y = append(y, twentySix[x[i]])
						flag = false
					}
				} else {
					y = append(y, twentySix[x[i]])
				}
			}
		} else {
			if i == len(x)-1 && byte(x[i]) == '1' {
				break
			}
			if byte(x[i]) == '1' {
				y = append(y, 'Z')
			} else {
				if byte(x[i]) == '0' {
					y = append(y, 'Y')
				} else {
					y = append(y, twentySix[x[i-1]])
					flag = false
				}
			}
		}
	}
	for i := len(y) - 1; i >= 0; i-- {
		fmt.Print(string(y[i]))
	}
	fmt.Println()
}

var tenToAny = map[int]byte{0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f', 16: 'g', 17: 'h', 18: 'i', 19: 'j', 20: 'k', 21: 'l', 22: 'm', 23: 'n', 24: 'o', 25: 'p', 26: 'q', 27: 'r', 28: 's', 29: 't', 30: 'u', 31: 'v', 32: 'w', 33: 'x', 34: 'y', 35: 'z', 36: ':', 37: ';', 38: '<', 39: '=', 40: '>', 41: '?', 42: '@', 43: '[', 44: ']', 45: '^', 46: '_', 47: '{', 48: '|', 49: '}', 50: 'A', 51: 'B', 52: 'C', 53: 'D', 54: 'E', 55: 'F', 56: 'G', 57: 'H', 58: 'I', 59: 'J', 60: 'K', 61: 'L', 62: 'M', 63: 'N', 64: 'O', 65: 'P', 66: 'Q', 67: 'R', 68: 'S', 69: 'T', 70: 'U', 71: 'V', 72: 'W', 73: 'X', 74: 'Y', 75: 'Z'}
var twentySix = map[byte]byte{'1': 'A', '2': 'B', '3': 'C', '4': 'D', '5': 'E', '6': 'F', '7': 'G', '8': 'H', '9': 'I', 'a': 'J', 'b': 'K', 'c': 'L', 'd': 'M', 'e': 'N', 'f': 'O', 'g': 'P', 'h': 'Q', 'i': 'R', 'j': 'S', 'k': 'T', 'l': 'U', 'm': 'V', 'n': 'W', 'o': 'X', 'p': 'Y', 'q': 'Z'}

func decimalToAny(num, n int) []byte {
	var newNumStr []byte
	var remainder int
	var remainderString byte
	for num != 0 {
		remainder = num % n
		remainderString = tenToAny[remainder]
		newNumStr = append(newNumStr, remainderString)

		num = num / n
	}
	return newNumStr
}
func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func TestThread(t *testing.T) {
	runtime.GOMAXPROCS(1) //First
	exit := make(chan int)
	go func() {
		close(exit)
		for {
			if true {
				println("Looping!") //Second
			}
		}
	}()
	<-exit
	println("Am I printed?")
}
func TestUpload(t *testing.T) {
	arra := []int64{1, 2, 3, 4}
	arrb := []string{"a", "b", "c", "d"}
	intChan := make(chan int64)
	strChan := make(chan string)
	var sg sync.WaitGroup
	sg.Add(1)
	go getInt(arra, intChan, strChan, &sg)
	go getStr(arrb, intChan, strChan)
	sg.Wait()
}

func getInt(intArr []int64, intChan chan int64, strChan chan string, sg *sync.WaitGroup) {
	for i := 0; i < len(intArr); i++ {
		fmt.Println(<-strChan)
		intChan <- intArr[i]
	}
	sg.Done()
}

func getStr(strArr []string, intChan chan int64, strChan chan string) {
	for i := 0; i < len(strArr); i++ {
		strChan <- strArr[i]
		fmt.Println(<-intChan)
	}
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
