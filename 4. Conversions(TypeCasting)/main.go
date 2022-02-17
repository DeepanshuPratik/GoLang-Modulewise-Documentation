package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("...Type Castings...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating of satisfaction / 10 : \n")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input)
	convertnum, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		panic(error)
	} else {
		fmt.Println("incremented rating : ", convertnum+3)
	}

}
