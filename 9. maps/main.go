package main

import "fmt"

func main() {
	fmt.Println("...Maps Construction...")
	coding_lang := make(map[string]string)

	coding_lang["cpp"] = "C++"
	coding_lang["go"] = "Golang"
	coding_lang["js"] = "Javascript"
	coding_lang["py"] = "Python"

	fmt.Println("maps coding_lang: ", coding_lang)
	fmt.Println("maps coding_lang extension to Fullname 'go': ", coding_lang["go"])
	delete(coding_lang, "cpp")
	fmt.Println("maps coding_lang cpp deleted: ", coding_lang)
	for key, val := range coding_lang {
		fmt.Printf("Key : %v  Value : %v \n", key, val)
	}

}
