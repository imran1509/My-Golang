package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//model for course - separate file

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

//middleware, helper - separate file

func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {

}

//controllers - separate file

func serveHome(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("<h1>Welcome to API by Imran</h1>"))
}

func getAllcourses(w http.ResponseWriter, r http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
