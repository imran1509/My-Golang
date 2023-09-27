package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("two")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
