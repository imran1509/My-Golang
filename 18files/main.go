package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to handling files in golang")
	content := "This need to go into a file - X.com/codenameimmy"

	file, err := os.Create("./MyTwitterHandle.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("lengt is: ", length)
	defer file.Close()
	readfile("./MyTwitterHandle.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}
