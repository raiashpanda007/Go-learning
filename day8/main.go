package main

import "fmt"

func main() {
	ashwin := User{"Ashwin", 20, "raiashwin005@gmail.com",true}
	fmt.Printf("The ashwin user info %+v\n",ashwin) // Using + adds property/method name too (ex. without Ashwin , 20 and with Name: Ashwin  )
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
