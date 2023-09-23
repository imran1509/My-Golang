package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance in golang; no super or parent concept

	imran := User{"Imran", "iimransaifi1509@gmail.com", true, 26}
	fmt.Println(imran)
	fmt.Printf("Imran's details are : %+v\n", imran)
	fmt.Printf("Name is %v and Email is %v\n", imran.Name, imran.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
