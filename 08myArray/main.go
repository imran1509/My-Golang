package main

import "fmt"

func main() {
	fmt.Println("welcome to Arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Grapes"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is: ", vegList)
}
