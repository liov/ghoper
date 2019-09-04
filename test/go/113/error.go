package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	fmt.Println(errors.As(bar(),&err))
}

func bar() error {
	return fmt.Errorf("%w",errors.New("UnWarp")).Unwrap()
}
