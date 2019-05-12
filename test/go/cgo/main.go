//+build go1.12

package main

//go:generate

//

//go run test/cgo

/*
#include <stdio.h>

void SayHelloInner(_GoString_ s);

void SayHelloExternal(const char* s);

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"
import "fmt"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
	C.SayHelloExternal(C.CString("Hello, World\n"))
	C.SayHelloInner("Hello, World\n")
}

//export SayHelloInner
func SayHelloInner(s string) {
	fmt.Print(s)
}
