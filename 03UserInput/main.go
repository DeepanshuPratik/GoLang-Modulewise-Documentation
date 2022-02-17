package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input lesson"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	// if dont want to store error use _ instead of variable
	input, _ := reader.ReadString('\n')
	fmt.Println("satisfaction rating : ", input)
}
