package main

import (
	"encoding/json"
	"fmt"
)

type coursename struct {
	Name     string  `json:"coursename"` //TO give alias to name 
	Price    int
	Platform string
	Password string `json:"-"` //To hide a field while creating json
	Tags     []string `json:"tags,omitempty"` // Inorder to make sure that if field is null than just don't throw that
}

func main() {
	EncodeJson();
}

func EncodeJson() {
	courses := []coursename{
		{
			Name:     "Go Mastery",
			Price:    4999,
			Platform: "Udemy",
			Password: "secure123",
			Tags:     []string{"golang", "backend", "concurrency"},
		},
		{
			Name:     "React - The Complete Guide",
			Price:    3999,
			Platform: "Coursera",
			Password: "reactrocks",
			Tags:     []string{"react", "frontend", "javascript"},
		},
		{
			Name:     "Full-Stack with TypeScript",
			Price:    6999,
			Platform: "Educative",
			Password: "typescriptftw",
			Tags:     []string{"typescript", "nodejs", "react", "fullstack"},
		},
		{
			Name:     "Microservices with gRPC in Go",
			Price:    5500,
			Platform: "Udemy",
			Password: "grpcguru",
			Tags:     []string{"golang", "grpc", "microservices", "backend"},
		},
		{
			Name:     "Building Scalable Apps with Kubernetes",
			Price:    8000,
			Platform: "Pluralsight",
			Password: "k8scloud",
			Tags:     nil,
		},
	}

	// finalJson, err := json.Marshal(courses)
	// For indentation :: 
	finalJson, err := json.MarshalIndent(courses, "" , "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("This is the json %s\n", finalJson);
}
