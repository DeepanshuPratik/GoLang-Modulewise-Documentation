package main

import "fmt"

func main() {

	fmt.Println("\t...LOOPS...")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Printf(days[i] + " ")
	}
	for i := range days {
		fmt.Println(days[i])
	}
	for _, day := range days {
		fmt.Printf("value is : %v \n", day)
	}

	// iteration type
	ranvalue := 1
	for ranvalue < 10 {
		if ranvalue == 5 {
			goto deep
		}
		fmt.Println(ranvalue)
		ranvalue++
	}
deep:
	fmt.Printf("Congo you got 5!")

}
