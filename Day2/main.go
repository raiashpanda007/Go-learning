package main

import "fmt"

const Hi string = "This is global hi" // Making captial letter making it PUBLIC

func main()  {
	var username string = "Ashwin Rai";
	fmt.Println(username);
	fmt.Printf("Variables type of username : %T\n", username)


	var isLoaggedIn bool = true;
	fmt.Println(isLoaggedIn);
	fmt.Printf("Variables type of isLoaggedIn : %T\n", isLoaggedIn)


	var anotherVariable float32
	fmt.Println(anotherVariable);
	fmt.Printf("Variables type of isLoaggedIn : %T\n", anotherVariable)


	// implictly typing 

	var website = "ashwin.projects.in";
	fmt.Println(website);

	// Declaring variable without var (can't be used outside main func)

	hello := "Hello my self Ashna";

	fmt.Println("hello is printed here", hello);

}

