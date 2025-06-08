package main

import "fmt"

// any or  interface {} or int | string | float32 | bool or comparable
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3, 4, 5})
	printSlice([]string{"apple", "banana", "cherry"})

	myStack := stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("Stack elements:", myStack.elements)
}
