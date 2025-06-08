package main

// For specific type
func variadicFunction(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// For any type
func anyVariadicFunction(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			println("Integer:", v)
		case string:
			println("String:", v)
		default:
			println("Unknown type")
		}
	}
}

func main() {

	result := variadicFunction(1, 2, 3, 4, 5)
	println("The sum is:", result)

	anyVariadicFunction(1, "hello", 3.14, "world", 42)

	nums := []int{10, 20, 30}
	variadicFunction(nums...)

}
