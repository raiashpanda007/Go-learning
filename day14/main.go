package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://api.github.com/users/raiashpanda007"

func main() {
	fmt.Println("Hitting github")
	data, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type :: %T\n", data);
	response ,err := io.ReadAll(data.Body)
	if(err != nil) {
		panic(err);
	}
	fmt.Println("The response from the github :: ",string(response))
	defer data.Body.Close()
	
	
}
