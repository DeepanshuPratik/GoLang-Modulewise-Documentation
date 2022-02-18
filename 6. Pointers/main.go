package main

import "fmt"

func main() {

	fmt.Println("....Pointers Concept....")
	mynumber := 69
	var ptr *int = &mynumber
	fmt.Println("Pointer value : ", *ptr, " pointer address : ", ptr)
	*ptr *= 4
	fmt.Println("Pointer value : ", mynumber)
}
