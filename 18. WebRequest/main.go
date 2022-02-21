package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("\t ... Web Request ...")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is: %T", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response is: ", string(databytes))
}
