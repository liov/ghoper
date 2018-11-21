package test

import (
	"fmt"
	"testing"
)

var a,b = 1,1



func TestUpload(t *testing.T)  {
	var str string = "1648464648"

	var data []byte = []byte(str)

	fmt.Println(data)

	b := string(data)

	fmt.Println(b)
}

