package main

import "fmt"

// Main get called by default and also you can not declare a func inside a function.

func adder (val1 int,val2 int) int {
	return val1 + val2
}
func proAdder (values ...int) int {
	total := 0;
	for _,val := range values {
		total += val
	}

	return total
}
func main() {
	res := adder(1,2)
	fmt.Println(res)
	prores := proAdder(1,2,3,4,5);
	fmt.Println("ProRes :: ", prores);
}