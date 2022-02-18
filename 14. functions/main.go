package main

import "fmt"

func main() {
	fmt.Println("\t ... Functions ...")
	greet()
	result := adder(35, 121)
	fmt.Println("Output of sum is : ", result)
	sum, messg := proAdder(2, 5, 7, 8, 14)
	fmt.Printf("Output of sum is : %v \n%v", sum, messg)

}

// fixed that 2 inputs should be there
func adder(a int, b int) int {
	return a + b
}

// no fixed no of inputs
func proAdder(values ...int) (int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum, "Hope we provide great service"
}

func greet() {
	fmt.Println("\t    NAMASTE INDIA!")
}
