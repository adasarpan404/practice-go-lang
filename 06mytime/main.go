package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:35 Monday"))

	createdDate := time.Date(2023, time.March, 10, 23, 23, 0, 34, time.Local)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
