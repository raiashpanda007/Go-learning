package main

import "fmt"

func main() {
	var Ashwin = User{"Ashwin", "raiashwin005@gmail.com", "Online", 10}
	Ashwin.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status string
	Age    int
}


// The object passed in the func is a copy or you can say by value.
func (u User) GetStatus() {
	fmt.Println("Is User Active", u.Status)
}




