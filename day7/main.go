package main

import "fmt"

func main() {
	languages := make(map[string]string) //Please remember for map and slices make is used 

	languages["ts"] = "typescript"
	languages["py"] = "python"
	languages["go"] = "golang"
	fmt.Println("List of all languages: ", languages)

	for key,value := range languages {
		fmt.Println("Key value pair of languages map", key , "\t", value , "\n");
	}
}
