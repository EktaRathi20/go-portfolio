package main

import "fmt"

func simpleArray() {
	var numbers [5]int

	fmt.Println(len(numbers))
}

func shortHandArray() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Short Hand Array:", numbers)
}

func twoDArray() {
	numbers := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Two Dimensional Array:", numbers)
}
func main() {
	simpleArray()
	shortHandArray()
	twoDArray()
}
