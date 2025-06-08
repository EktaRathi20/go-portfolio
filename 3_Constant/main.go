package main

import "fmt"

func main() {
	const pi float32 = 3.14
	// pi = 3.14159 // This will cause a compile-time error because pi is a constant and cannot be reassigned
	fmt.Println("Value of pi is:", pi)

	// way to declare a constant
	const (
		day   = "Monday"
		month = "January"
		year  = 2023
	)

	fmt.Println("Today is", day, month, year)
}
