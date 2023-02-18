package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.September, 15, 05, 05, 0, 0, time.UTC) //Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
