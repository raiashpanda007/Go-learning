package main

import "fmt"

func main() {
	fmt.Println("Loops for weekdays")
	var days = []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday"}
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("Days :: ",days[d])
	// }

	// for d := range(days) {
	// 	fmt.Println("Days :: ",days[d])
	// }

	for _,day := range days {
		fmt.Println("Days :: ",day)
	}

}
