package main

import "fmt"

const prefix = "Hello, "

func functionWithoutParams() string {
	return "Hello, world"
}

func functionWithParams(name string) string {
	return "Hello, " + name
}

func functionWithMultipleParams(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return prefix + name
}

func functionWithVariable() string {
	return prefix + "world"
}

// IF and ELSE
func functionWithIfElse(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	if language == "French" {
		return "Holaa, " + name
	}
	return "Hello" + name
}

func styling() {
	name := "Amit"
	age := 25

	fmt.Println("My name is", name, "and my age is", age)

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

}

func storeFormattedString(name string, age int) string {
	return fmt.Sprintf("My name is %s and I am %d years old.", name, age)
}

func getNameAndAge() (string, int) {
	return "Ravi", 25
}

func getLanguage() (string, string, string) {
	return "Go", "js", "python"
}

func functionWithFunctionParameters(fn func(a int) int) {
	result := fn(5)
	fmt.Println("Result from function with function parameter:", result)
}

// Entry point of the program
func main() {
	fmt.Println(functionWithoutParams())
	fmt.Println(functionWithParams("World"))
	fmt.Println(functionWithMultipleParams("Test", "Spanish"))
	fmt.Println(functionWithVariable())
	fmt.Println(functionWithIfElse("Test", "Spanish"))
	styling()
	message := storeFormattedString("Ekta", 23)
	fmt.Println(message)

	// If we want to ignore the name returned by getNameAndAge, we can use the blank identifier _
	_, age := getNameAndAge()
	fmt.Println(age)

	l1, l2, l3 := getLanguage()
	fmt.Println(l1, l2, l3)

	fn := func(a int) int {
		return a * 2
	}

	functionWithFunctionParameters(fn)
}
