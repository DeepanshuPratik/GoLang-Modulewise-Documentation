package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("\t ...Switch Case...")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // as 6 is exclusive

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1 , can open up your dice or move")
	case 2:
		fmt.Println("you can move 2 spots ahead")
	case 3:
		fmt.Println("you can move 3 spots ahead")
		fallthrough // it will fall into 4 after executing 3
	case 4:
		fmt.Println("you can move 4 spots ahead")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spots ahead")
	case 6:
		fmt.Println("Value is 1 , can open up your dice or move")

	default:
		fmt.Println("What was that!")
	}

}
