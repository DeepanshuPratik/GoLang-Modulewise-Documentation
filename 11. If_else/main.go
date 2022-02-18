package main

import "fmt"

func main() {
	fmt.Println("\t...If else conditionings...")
	var message string
	// initiallization and conditioning done in single if .
	if countCust := 35; countCust > 25 {
		message = "Good & Profitable customer"
	} else {
		message = "Less priority Customer"
	}
	fmt.Println(message)

}
