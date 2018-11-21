package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	a:=1
	b,_ :=json.Marshal(a)
	/*	switch b.(type) {
		case int:
			fmt.Println("true")
		}*/
	fmt.Print(b,)
}
