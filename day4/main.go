package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Please buy this stock");
	reader := bufio.NewReader(os.Stdin);

	input ,_ := reader.ReadString('\n');
	fmt.Println("Buying this stock", input);
	numRating,err := strconv.ParseFloat(strings.Trim(input, "\n"), 64);

	if err != nil {
		fmt.Println("Error",err);
	} else {
		fmt.Println("numRating",numRating + 1);
	}


}
