package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("...Creating Structs...")
	// no inheritance and no super/parent class
	Deepanshu := User{"Deepanshu Pratik", "deepanshu.pratik@gmail.com", true, 20}

	fmt.Printf("Details of deepanshu: %v\n", Deepanshu)
	fmt.Printf("Details of deepanshu: %+v\n", Deepanshu)
	fmt.Printf("Details of deepanshu Name: %+v\n", Deepanshu.Name)
	fmt.Printf("Details of deepanshu Email: %v\n", Deepanshu.Email)
}
