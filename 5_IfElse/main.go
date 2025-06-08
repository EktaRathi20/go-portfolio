package main

import "fmt"

func main() {
	var name string
	var age int
	var yourPosition string

	fmt.Print("Enter your name and age: ")
	fmt.Scanln(&name, &age)

	if age < 13 {
		yourPosition = "You're a child"
	} else if age < 20 {
		yourPosition = "You're a teenager"
	} else if age < 60 {
		yourPosition = "You're a adult"
	} else {
		yourPosition = "You're a senior citizen"
	}

	fmt.Printf("Hi %s, you are %d years old. %s.", name, age, yourPosition)

	// we can declare variable inside the if statement
	if age := 18; age < 20 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are an adult.")
	}

	// go does not support ternary operator like other languages

}
