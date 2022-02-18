package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("\t ...Files...")
	content := "I am a noob following some great content and in path to become mediocre!"

	file, err := os.Create("./mysampletxt.txt")

	// if err != nil {
	// 	panic(err)
	// }
	errorChecker(err)

	length, err := io.WriteString(file, content)

	errorChecker(err)

	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./mysampletxt.txt")
}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	errorChecker(err)
	fmt.Println("Data stored is : \n", string(databytes))
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
