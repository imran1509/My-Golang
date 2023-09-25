package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
	case 4:
		fmt.Println("you can move 4 spots")
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move to 6 spot and roll the dice again")
	default:
		fmt.Println("What was that!")
	}
}
