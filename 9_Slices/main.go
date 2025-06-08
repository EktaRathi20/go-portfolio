package main

import (
	"fmt"
	"slices"
)

func nilSlices() {
	var numbers []int
	fmt.Println((numbers == nil))
}

func simpleSlices() {
	// 1. which type of slice to create
	// 2. how many elements to create
	// 3. what is the capacity of the slice (if we want to set it manually)
	var number1 = make([]int, 5)
	fmt.Println("Capacity of slice:", cap(number1))
	fmt.Println("Length of slice:", len(number1))

	var number2 = make([]int, 5, 10)
	fmt.Println("Capacity of slice:", cap(number2))
	fmt.Println("Length of slice:", len(number2))
}

func appendElementToSlice() {
	var numbers = make([]int, 1)
	numbers = append(numbers, 100)
	numbers = append(numbers, 200)

	fmt.Println("Slice after appending elements:", numbers)
}

// in this it will double the capacity of the slice when it reaches its limit
func autoResizeSlice() {
	var numbers = make([]int, 0, 5)
	fmt.Println("Initial capacity:", cap(numbers))

	for i := 0; i < 6; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println("Final slice:", numbers)
	fmt.Println("Final length:", len(numbers))
	fmt.Println("Final capacity:", cap(numbers))
}

func shortHandSlice() {
	numbers := []int{}
	numbers = append(numbers, 10)
	fmt.Println("Short Hand Slice:", numbers)
}

func copySlice() {
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println("Source Slice:", source)
	fmt.Println("Destination Slice after copy:", destination)
}

func sliceSlice() {
	numbers := []int{1, 2, 3, 4, 5}

	// include index 0 and exclude index 4
	sliced := numbers[0:4]
	// OR sliced := numbers[:4]

	sliced2 := numbers[1:]

	fmt.Println("Original Slice:", numbers)
	fmt.Println("Sliced Slice:", sliced)
	fmt.Println("Sliced2 Slice:", sliced2)
}

func areSlicesEqual() {
	slice1 := []int{1, 2}
	slice2 := []int{1, 2}

	isEqual := slices.Equal(slice1, slice2)
	fmt.Println("Are slices equal?", isEqual)
}

func twoDSlices() {
	numbers := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Two Dimensional Slice:", numbers)
}

func main() {
	nilSlices()
	simpleSlices()
	appendElementToSlice()
	autoResizeSlice()
	shortHandSlice()
	copySlice()
	sliceSlice()
	areSlicesEqual()
	twoDSlices()
}
