package main

import "fmt"
import "sort"

func main() {
	fmt.Println("This day is about learning array and slices")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))
	fmt.Println("----------------------------------------------------------------------------------Slices----------------------------------------------------------")
	var fruitSlice = []string{}
	fruitSlice = append(fruitSlice, "Seb", "tamatar", "peach")
	fmt.Println("After appending the slices", fruitSlice)

	// by using make dynamic memory allocation

	var highScores = make([]int, 4)
	highScores[0] = 100
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 this is not possible  but  can add by appending into it
	highScores = append(highScores, 555, 666, 321)
	sort.Ints(highScores)

	fmt.Println("--------------------------------------------------------Delete index------------------------------------------------------------")

	var courses = []string{"reactjs", "js", "ts", "python"};

	index := 1;
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("Print courses" , courses)

}
