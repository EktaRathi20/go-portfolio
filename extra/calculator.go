package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) float64 {
	return float64(a) / float64(b)
}

func calculator() {
	var choice int
	var a, b int

	for {
		fmt.Println("\n--- Calculator ---")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Exiting... Bye!")
			break
		}

		fmt.Print("Enter two numbers: ")
		fmt.Scanln(&a, &b)

		switch choice {
		case 1:
			fmt.Println("Result:", add(a, b))
		case 2:
			fmt.Println("Result:", subtract(a, b))
		case 3:
			fmt.Println("Result:", multiply(a, b))
		case 4:
			fmt.Println("Result:", divide(a, b))
		default:
			fmt.Println("Invalid choice")
		}
	}
}

