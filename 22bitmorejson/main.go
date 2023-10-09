package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeonline.com", "imr123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
