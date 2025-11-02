package main

import (
	"fmt"
	"io"
	"os"

)

func main() {
	content := "This is the file content by Ashwin rai"

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is", length);

	defer file.Close()

	readFile("./test.txt");

}

func readFile (fileName string) {
	data, err := os.ReadFile(fileName); // Gotcha that the fileData that is read here is in byte format;

	if(err != nil) {
		panic(err);
	}
	fmt.Println("Raw data",string(data));

}
