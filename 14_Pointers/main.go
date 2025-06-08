package main

import "fmt"

func callByValue(x int) {
	x = 10
	fmt.Println("Inside callByValue, x =", x)
}

func callByReference(y *int) {
	*y = 10 // Dereference the pointer to change the value
	fmt.Println("Inside callByReference, y =", *y)
}

func main() {
	x := 5
	fmt.Println("Before callByValue, x =", x)
	callByValue(x)
	fmt.Println("After callByValue, x =", x)

	y := 5
	fmt.Println("Before callByReference, y =", y)
	callByReference(&y) // Pass the address of y to the function
	fmt.Println("After callByReference, y =", y)
}
