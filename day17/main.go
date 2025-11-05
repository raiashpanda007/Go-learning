package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"` //TO give alias to name
	Price    int
	Platform string
	Password string   `json:"-"`              //To hide a field while creating json
	Tags     []string `json:"tags,omitempty"` // Inorder to make sure that if field is null than just don't throw that
}

func main() {

	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "Microservices with gRPC in Go",
                "Price": 5500,
                "Platform": "Udemy",
                "tags": [
                        "golang",
                        "grpc",
                        "microservices",
                        "backend"
                ]
    }`)

	var myCourse courses

	isValid := json.Valid(jsonDataFromWeb)
	if isValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Printf("Invalid json ");
	}

}
