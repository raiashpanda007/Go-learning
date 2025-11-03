package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	
)

const URL = "http://localhost:3000/post"

func main() {
	PostJSONRequest()
}

func PostJSONRequest() {
	requestBody := strings.NewReader(`
		{
			"message":"Ashwin is my name"
		}
	`);
	response , err := http.Post(URL, "application/json",requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body);

	if err != nil {
		panic(err)
	}

	fmt.Println("data :: ", string(data));
}
