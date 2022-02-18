package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("...Slices Topic to be covered...")
	var ToppersList = []string{}

	fmt.Printf(" Type:%T\n ", ToppersList)
	ToppersList = append(ToppersList, "Vikalp Tripathi", "Divyansh Kushwaha", "Soumya Ranjan", "Kshitiz Kumar Singh")
	fmt.Println(" List : \n ", ToppersList)
	// index 1 to 2, 3 is exclusive
	fmt.Println(" List : \n ", ToppersList[1:3])
	// index starting till 1 as 2 is exclusive
	fmt.Println(" List : \n ", ToppersList[:2])
	// index 2 till end
	fmt.Println(" List : \n ", ToppersList[2:])

	// Make using in slices
	highscores := make([]int, 4)

	highscores[0] = 97
	highscores[1] = 94
	highscores[2] = 92
	highscores[3] = 89

	fmt.Printf(" Type:%T\n ", highscores)
	fmt.Println("Highscores: ", highscores)

	highscores = append(highscores, 99, 90, 87, 85)
	fmt.Printf(" Type:%T\n ", highscores)
	fmt.Println("Highscores: ", highscores)
	fmt.Println("Are highscores sorted: ", sort.IntsAreSorted(highscores))

	// sorting slice
	sort.Ints(highscores)
	fmt.Println("Highscores: ", highscores)
	fmt.Println("Are highscores sorted: ", sort.IntsAreSorted(highscores))

	// how to remove element index based in slices
	var courses = []string{"react.js", "typescript", "Golang", "python", "C++", "Kotlin"}
	fmt.Println("\n\nCourses : ", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("\nCourses deleting 2: ", courses)

}
