package main

import "fmt"

// for -> only contruct that can be used to iterate over a range of values

func whileLoopExample() {
	i := 0
	for i < 5 {
		fmt.Println("Iteration:", i)
		i++
	}
}

// func infiniteLoopExample() {
// 	for {
// 		fmt.Println("This is an infinite loop")
// 		// Uncomment the next line to break the loop
// 		// break
// 	}
// }

func forLoopExample() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
}

// break and continue statements can be used to control the flow of loops
// break exits the loop, while continue skips to the next iteration
// example of break and continue:
func breakAndContinueExample() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at i =", i)
			break // exits the loop when i is 5
		}
		if i%2 == 0 {
			fmt.Println("Skipping even number:", i)
			continue // skips the rest of the loop for even numbers
		}
		fmt.Println("Odd number:", i)
	}
}

// range
func rangeExample() {
	for i := range 3 {
		fmt.Println("Iteration:", i)
	}
}

func main() {
	whileLoopExample()
	// infiniteLoopExample()
	forLoopExample()
	breakAndContinueExample()
	rangeExample()
}
