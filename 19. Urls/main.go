package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://deepanshu.dev:8080/learn?courses=golang&paymentid=fvewsf425sdfe32eesf"

func main() {
	fmt.Println("\t ...URL Handelling ...")
	fmt.Println(myurl)

	content, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(content.Scheme)
	fmt.Println(content.RawQuery)
	fmt.Println(content.Path)
	fmt.Println(content.Hostname())
	fmt.Println(content.Host)

	params := content.Query()
	fmt.Printf("THE TYPE OF PARAM IS : %T", params)
	fmt.Println("\n", params)
	for _, val := range params {
		fmt.Println(val)
	}
}
