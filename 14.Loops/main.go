package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in Golang")

	Days := []string{"Sunday", "Tusesday", "wednesday", "Friday", "saturday"}
	fmt.Println(Days)

	for d := 0; d < len(Days); d++ {
		fmt.Println(Days[d])
	}

	for i := range Days {
		fmt.Println(Days[i])
	}

	for index, Day := range Days {
		fmt.Printf("index is %v and value is %v\n", index, Day)
	}

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 5 {
			break
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}
	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}
}
