package main

import "fmt"

func rangeWithSlice() {
	s := []int{1, 2, 3, 4, 5}
	for index, num := range s {
		fmt.Printf("Index: %d, Value: %d\n", index, num)
	}
}

func rangeWithMap() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func rangeWithString() {
	str := "GoLang"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}

func main() {
	rangeWithSlice()
	rangeWithMap()
	rangeWithString()
}
