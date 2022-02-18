package main

import "fmt"

func main() {
	fmt.Println("...Slices Topic to be covered...")
	var ToppersList = []string{}

	fmt.Printf(" Type:%T\n ", ToppersList)
	ToppersList = append(ToppersList, "Vikalp Tripathi", "Divyansh Kushwaha", "Soumya Ranjan", "Kshitiz Kumar Singh")
	fmt.Println(" List : \n ", ToppersList)
	fmt.Println(" List : \n ", ToppersList[1:3])
	fmt.Println(" List : \n ", ToppersList[:2])
	fmt.Println(" List : \n ", ToppersList[2:])
}
