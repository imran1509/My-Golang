package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hjfhd895"

func main() {
	fmt.Println(myurl)

	//parsing an Url using url library
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	//properties and methods of a parsed url
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//This gives queries in a more structured format
	qparams := result.Query()

	//The type of qparams turns out to be url.values which is just a
	//fancy name for Key-value pairs
	//E.g. coursename is key and value is reactjs
	fmt.Printf("The type of query parameters is: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for index, value := range qparams {
		fmt.Println(index, value)
	}

	//Constructing a URL using difft params
	//Always pass it using reference i.e. &
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
