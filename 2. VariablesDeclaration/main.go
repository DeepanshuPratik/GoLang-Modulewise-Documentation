package main

import "fmt"

func main() {

	var username string = "Deepanshu_Pratik"
	fmt.Println(username)
	fmt.Printf("The variable is of type: %T \n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of type: %T \n", isLoggedIn)
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("The variable is of type: %T \n", smallval)
	var slr complex128
	fmt.Println(slr)

	var app = "SOCIAL IIITNR"
	fmt.Println(app)

	roll_no := 201000217
	fmt.Println(roll_no)
}

// If variable has first letter caps it implies its public
