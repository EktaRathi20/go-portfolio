package main

import (
	"fmt"
	"maps"
)

func createMap() {
	m := make(map[string]string)
	fmt.Println(m)
}

func addKeyValuePair() {
	// 1. create a map
	m := make(map[string]string)

	// 2. add key-value pairs
	m["name"] = "go-lang"

	// 3. access a value by key
	fmt.Println(m["name"])
}

func accessNonExistentKey() {
	// 1. create a map
	m := make(map[string]string)

	// 2. access a value by key that doesn't exist
	// This will return the zero value for the value type, which is an empty string for strings
	fmt.Println(m["age"])
}

func mapLength() {
	// 1. create a map
	m := make(map[string]string)

	// 2. add key-value pairs
	m["name"] = "go-lang"

	// 3. Length of the map
	fmt.Println("Length of map:", len(m))
}
func deleteKeyValuePair() {
	// 1. create a map
	m := make(map[string]string)

	// 2. add key-value pairs
	m["name"] = "go-lang"

	// 3. delete a key-value pair
	delete(m, "name")

	// 4. print the map after deletion
	fmt.Println("Map after deletion:", m)
}

func shorthandMapCreation() {
	// 1. create a map with initial values
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// 2. print the map
	fmt.Println(m)
}

func checkKeyExists() {
	// 1. create a map with initial values
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// 2. check if a key exists in the map
	value, exists := m["two"]

	if exists {
		fmt.Println("Key exists in the map")
		fmt.Println("Value for 'two':", value)
	} else {
		fmt.Println("Key does not exist in the map")
	}
}

func areMapsEqual() {
	m1 := map[string]int{"one": 1, "two": 2}
	m2 := map[string]int{"one": 1, "two": 2}

	fmt.Println("Are maps equal?", maps.Equal(m1, m2))
}
func main() {
	createMap()
	addKeyValuePair()
	accessNonExistentKey()
	mapLength()
	deleteKeyValuePair()
	shorthandMapCreation()
	checkKeyExists()
	areMapsEqual()
}
