package main

import "fmt"

func main() {
	fmt.Println("\t ... DEFERS ...")
	defer fmt.Printf("\nam I at last ?")       // st 1
	defer fmt.Printf("\nAm I at secondlast ?") // st 2
	defer fmt.Printf("\n Am I at thirdlast ?") // st 3
	// In defers FILO works, as st 1 enters first => it will print at last ...
	Defer() // st 4
}
func Defer() {
	for i := 1; i < 10; i++ {
		defer fmt.Printf("I am no %v\n", i)
	}
}
