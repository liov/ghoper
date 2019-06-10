package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func init()  {
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
}


func main()  {
	time.Sleep(*period)
	fmt.Println("end")
}
