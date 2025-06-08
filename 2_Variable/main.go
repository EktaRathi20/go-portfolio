package main

import "fmt"

func main() {
	// Way 1 to declare a variable
	var name1 string = "Hello, world 1"

	// Way 2 to declare a variable
	var name2 = "Hello, world 2"

	// Way 3 to declare a variable Shorthand Syntax 
	// (only works inside functions)
	name3 := "Hello, world 3"

	fmt.Println(name1, name2, name3)
}
