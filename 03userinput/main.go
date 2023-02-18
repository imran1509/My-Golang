package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok syntax || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("type of the rating is %T", input)
}
