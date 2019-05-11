package main

//go run test/cgo

/*
#include <stdio.h>

void SayHelloExternal(const char* s);

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
	C.SayHelloExternal(C.CString("Hello, World\n"))
}
