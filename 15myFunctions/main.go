package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(6, 9)

	fmt.Println("results is : ", result)

	proRes, myMessage := proAdder(1, 3, 4, 5, 6, 8, 9)
	fmt.Println("Pro resut is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Hello from Golang")
}
