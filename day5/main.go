package main

import "fmt"

func main(){
	myNumber := 23;
	var ptr = &myNumber;
	fmt.Println("Value of address  actual pointer is ",ptr);
	fmt.Println("Value of pointer", * ptr);

}
