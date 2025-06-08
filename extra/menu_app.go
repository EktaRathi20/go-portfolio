package main

import "fmt"

func menuApp() {
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Say Hello")
		fmt.Println("2. Add Two Numbers")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("Hello there! 👋")
		} else if choice == 2 {
			var a, b int
			fmt.Print("Enter two numbers: ")
			fmt.Scanln(&a, &b)
			fmt.Println("Sum is:", a+b)
		} else if choice == 3 {
			fmt.Println("Exiting... Bye! 👋")
			break
		} else {
			fmt.Println("Invalid choice. Try again.")
		}

	}
}
