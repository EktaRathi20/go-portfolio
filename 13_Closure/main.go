package main

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}
func main() {

	counterFunc := counter()

	for i := 0; i < 10; i++ {
		println(counterFunc())
	}

}
