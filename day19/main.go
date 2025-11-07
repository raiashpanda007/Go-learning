package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
)

type Courses struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"courseName"`
	Price      int     `json:"int"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string
}

var courses []Courses

// Crating a fake middleware
func (c *Courses) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses);
   
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", GetCourses).Methods("GET")

    fmt.Println("Server started at :4000")
    log.Fatal(http.ListenAndServe(":4000", router))
}
