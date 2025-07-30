package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Please provide input \n");
	reader := bufio.NewReader(os.Stdin);
	

	// Go don't have try catch so it uses comma , ok thing

	input,_ := reader.ReadString('\n');
	fmt.Println("You typed this :: ",input);
}