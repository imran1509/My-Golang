package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages : ", languages)
	fmt.Println("Js shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages : ", languages)

	//loops are interesting in Golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
