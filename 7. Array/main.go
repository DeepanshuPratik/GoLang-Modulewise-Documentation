package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("...Arrays Concept...")
	fmt.Println("Enter no of toppers : ")
	count, _ := reader.ReadString('\n')
	var ToppersList [7]string
	ToppersList[0] = "Vikalp Tripathi"
	ToppersList[1] = "Divyansh Kushwaha"
	ToppersList[2] = "Soumya Ranjan"
	ToppersList[4] = "Kshitiz Kumar Singh"

	fmt.Println("COunt: ", count, " ToppersList: ", ToppersList, " Length: ", len(ToppersList))

}
