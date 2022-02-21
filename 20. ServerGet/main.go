package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("\t ...Server Get...")
	//PerformGetRequest()
	PerformPOstJsonRequest()
}
func PerformGetRequest() {
	const weburl = "http://localhost:8000/get"
	res, err := http.Get(weburl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("res type is %T \nits status code is %v\nContent Length : %v", res, res.StatusCode, res.ContentLength)
	var respString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	byteCount, _ := respString.Write(content)
	fmt.Println("\nBYTE LENGTH : ", byteCount)
	//fmt.Println("\ncontent is : ", string(content))
	//Another way
	fmt.Println("\n", respString.String())

}
func PerformPOstJsonRequest() {
	const weburl = "http://localhost:8000/post"

	// fake payload
	contentPushed := strings.NewReader(`
		{
			"coursename":"GoLang",
			"authorname":"Deepanshu",
			"price":"$200"
		}
	`)
	resp, err := http.Post(weburl, "application/json", contentPushed)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("content is : ", string(content))
}
